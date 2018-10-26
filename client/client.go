package ankadbclient

import (
	"context"

	pb "github.com/zhs007/ankadb/proto"
	"google.golang.org/grpc"
)

// AnkaClient -
type AnkaClient interface {
	Start(addr string) error
	Stop() error
	Query(ctx context.Context, request string, varval string) (*pb.ReplyQuery, error)
}

// AnkaClient -
type ankaClient struct {
	addr   string
	conn   *grpc.ClientConn
	client pb.AnkaDBServClient
}

// NewClient -
func NewClient() AnkaClient {
	return &ankaClient{}
}

func (c *ankaClient) Start(addr string) error {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return err
	}

	c.conn = conn
	c.client = pb.NewAnkaDBServClient(conn)

	return nil
}

func (c *ankaClient) Stop() error {
	c.addr = ""
	c.conn = nil

	return nil
}

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
