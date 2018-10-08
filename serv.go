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
	chanServ chan int
}

// newServer -
func newServer(anka *AnkaDB) (*ankaServer, error) {
	lis, err := net.Listen("tcp", anka.cfg.AddrBind)
	if err != nil {
		return nil, err
	}

	// log.Info("Listen", zap.String("addr", node.myinfo.BindAddr))

	grpcServ := grpc.NewServer()
	s := &ankaServer{
		anka:     anka,
		lis:      lis,
		grpcServ: grpcServ,
		chanServ: make(chan int),
	}

	pb.RegisterAnkaDBServServer(grpcServ, s)

	return s, nil
}

func (s *ankaServer) start() (err error) {
	// fmt.Print("start...")
	err = s.grpcServ.Serve(s.lis)
	// fmt.Print("end start...")
	s.chanServ <- 0
	// fmt.Print("exit")

	return
}

func (s *ankaServer) stop() {
	// fmt.Print("stop0...")
	s.lis.Close()
	// fmt.Print("stop1...")
	s.grpcServ.Stop()
	// fmt.Print("stop2...")

	return
}

// Query implements ankadbpb.ankaServer
func (s *ankaServer) Query(ctx context.Context, in *pb.Query) (*pb.ReplyQuery, error) {
	var mapval map[string]interface{}
	err := json.Unmarshal([]byte(in.GetVarData()), &mapval)
	if err != nil {
		rq := pb.ReplyQuery{
			Code: pb.CODE_VAR_PARSE_ERR,
			Err:  err.Error(),
		}
		return &rq, nil
	}

	result, err := s.anka.logic.OnQuery(in.GetQueryData(), mapval)
	if err != nil {
		rq := pb.ReplyQuery{
			Code: pb.CODE_LOGIC_ONQUERY_ERR,
			Err:  err.Error(),
		}
		return &rq, nil
	}

	buf, _ := json.Marshal(result)

	return &pb.ReplyQuery{
		Code:   0,
		Result: string(buf),
	}, nil
}
