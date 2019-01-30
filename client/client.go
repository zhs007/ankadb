package ankadbclient

import (
	"context"

	pb "github.com/zhs007/ankadb/proto"
	"google.golang.org/grpc"
)

// AnkaClient -
type AnkaClient interface {
	// Start - start a client
	Start(addr string) error
	// Stop - stop a client
	Stop() error

	// Query - query a GraphQL
	Query(ctx context.Context, request string, varval string) (*pb.ReplyQuery, error)

	// Get - get value with the key
	Get(ctx context.Context, dbname string, key string) (*pb.ReplyGetValue, error)
	// Set - set value with the key
	Set(ctx context.Context, dbname string, key string, value []byte) (*pb.ReplySetValue, error)
}

// AnkaClient -
type ankaClient struct {
	addr   string
	conn   *grpc.ClientConn
	client pb.AnkaDBServClient
}

// NewClient - new a client
func NewClient() AnkaClient {
	return &ankaClient{}
}

// Start - start a client
func (c *ankaClient) Start(addr string) error {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return err
	}

	c.conn = conn
	c.client = pb.NewAnkaDBServClient(conn)

	return nil
}

// Stop - stop a client
func (c *ankaClient) Stop() error {
	c.addr = ""
	c.conn = nil

	return nil
}

// Query - query a GraphQL
func (c *ankaClient) Query(ctx context.Context, request string, varval string) (*pb.ReplyQuery, error) {
	if c.conn == nil {
		return nil, ErrNoConn
	}

	curctx, cancel := context.WithCancel(ctx)
	defer cancel()

	r, err := c.client.Query(curctx, &pb.Query{
		// Name:      name,
		QueryData: request,
		VarData:   varval,
	})
	if err != nil {
		return nil, err
	}

	if r.Err == "" {
		return r, nil
	}

	return r, nil
}

// Get - get value with the key
func (c *ankaClient) Get(ctx context.Context, dbname string, key string) (*pb.ReplyGetValue, error) {
	if c.conn == nil {
		return nil, ErrNoConn
	}

	curctx, cancel := context.WithCancel(ctx)
	defer cancel()

	r, err := c.client.Get(curctx, &pb.GetValue{
		NameDB: dbname,
		Key:    key,
	})
	if err != nil {
		return nil, err
	}

	if r.Err == "" {
		return r, nil
	}

	return r, nil
}

// Set - set value with the key
func (c *ankaClient) Set(ctx context.Context, dbname string, key string, value []byte) (*pb.ReplySetValue, error) {
	if c.conn == nil {
		return nil, ErrNoConn
	}

	curctx, cancel := context.WithCancel(ctx)
	defer cancel()

	r, err := c.client.Set(curctx, &pb.SetValue{
		NameDB: dbname,
		Key:    key,
		Value:  value,
	})
	if err != nil {
		return nil, err
	}

	if r.Err == "" {
		return r, nil
	}

	return r, nil
}
