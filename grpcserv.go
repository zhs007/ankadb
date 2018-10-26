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
	anka     *AnkaDB
	lis      net.Listener
	grpcServ *grpc.Server
	// chanServ chan int
}

// newServer -
func newServer(anka *AnkaDB) (*ankaServer, error) {
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
		// chanServ: make(chan int),
	}

	pb.RegisterAnkaDBServServer(grpcServ, s)

	return s, nil
}

func (s *ankaServer) start(ctx context.Context) (err error) {
	err = s.grpcServ.Serve(s.lis)

	// s.chanServ <- 0

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
