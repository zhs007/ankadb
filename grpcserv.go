package ankadb

import (
	"context"
	"encoding/json"
	"net"

	pb "github.com/zhs007/ankadb/proto"
	"google.golang.org/grpc"
)

// ankaServer
type ankaServer struct {
	anka     *ankaDB
	lis      net.Listener
	grpcServ *grpc.Server
}

// newServer -
func newServer(anka *ankaDB) (*ankaServer, error) {
	lis, err := net.Listen("tcp", anka.cfg.AddrGRPC)
	if err != nil {
		return nil, ErrGRPCListen
	}

	// log.Info("Listen", zap.String("addr", node.myinfo.BindAddr))

	grpcServ := grpc.NewServer()
	s := &ankaServer{
		anka:     anka,
		lis:      lis,
		grpcServ: grpcServ,
	}

	pb.RegisterAnkaDBServServer(grpcServ, s)

	return s, nil
}

func (s *ankaServer) start(ctx context.Context) (err error) {
	err = s.grpcServ.Serve(s.lis)

	return
}

func (s *ankaServer) stop() {
	s.lis.Close()

	s.grpcServ.Stop()

	return
}

// Query implements ankadbpb.ankaServer
func (s *ankaServer) Query(ctx context.Context, in *pb.Query) (*pb.ReplyQuery, error) {
	var mapval map[string]interface{}
	err := json.Unmarshal([]byte(in.GetVarData()), &mapval)
	if err != nil {
		rq := pb.ReplyQuery{
			Err: err.Error(),
		}
		return &rq, nil
	}

	// curdb := s.anka.MgrDB.GetDB(in.GetName())
	curctx := context.WithValue(ctx, interface{}("ankadb"), s.anka)

	result, err := s.anka.logic.OnQuery(curctx, in.GetQueryData(), mapval)
	if err != nil {
		rq := pb.ReplyQuery{
			Err: err.Error(),
		}
		return &rq, nil
	}

	buf, _ := json.Marshal(result)

	return &pb.ReplyQuery{
		Result: string(buf),
	}, nil
}

// QueryStream implements ankadbpb.ankaServer
func (s *ankaServer) QueryStream(in *pb.Query, gs pb.AnkaDBServ_QueryStreamServer) error {
	var mapval map[string]interface{}
	err := json.Unmarshal([]byte(in.GetVarData()), &mapval)
	if err != nil {
		gs.Send(&pb.ReplyQuery{
			Err: err.Error(),
		})

		return nil
	}

	err = s.anka.logic.OnQueryStream(gs.Context(), in.GetQueryData(), mapval, func(rq *pb.ReplyQuery) {
		gs.Send(rq)
	})
	if err != nil {
		gs.Send(&pb.ReplyQuery{
			Err: err.Error(),
		})
	}

	return nil
}

// Get implements ankadbpb.ankaServer
func (s *ankaServer) Get(ctx context.Context, in *pb.GetValue) (*pb.ReplyGetValue, error) {
	buf, err := s.anka.Get(ctx, in.NameDB, in.Key)
	if err != nil {
		return &pb.ReplyGetValue{
			Err: err.Error(),
		}, err
	}

	return &pb.ReplyGetValue{
		Value: buf,
	}, nil
}

// Set implements ankadbpb.ankaServer
func (s *ankaServer) Set(ctx context.Context, in *pb.SetValue) (*pb.ReplySetValue, error) {
	err := s.anka.Set(ctx, in.NameDB, in.Key, in.Value)
	if err != nil {
		return &pb.ReplySetValue{
			Err: err.Error(),
		}, err
	}

	return &pb.ReplySetValue{}, nil
}

// SetQueryTemplate implements ankadbpb.ankaServer
func (s *ankaServer) SetQueryTemplate(ctx context.Context, in *pb.QueryTemplate) (*pb.ReplyQueryTemplate, error) {
	err := s.anka.SetQueryTemplate(in.QueryTemplateName, in.QueryData)
	if err != nil {
		return &pb.ReplyQueryTemplate{
			Err: err.Error(),
		}, err
	}

	return &pb.ReplyQueryTemplate{}, nil
}
