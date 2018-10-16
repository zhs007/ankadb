package ankadb

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/zhs007/ankadb/err"
	pb "github.com/zhs007/ankadb/proto"
)

// ankaHTTPServer
type ankaHTTPServer struct {
	anka     *AnkaDB
	lis      net.Listener
	chanServ chan int
}

func (s *ankaHTTPServer) procGraphQL(w http.ResponseWriter, r *http.Request) *graphql.Result {
	// ankadbname := r.Header.Get("Ankadbname")

	req, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	var mapreq map[string]interface{}
	err := json.Unmarshal([]byte(req), &mapreq)
	if err != nil {
		return MakeGraphQLErrorResult(pb.CODE_HTTP_BODY_PARSE_ERR)
	}

	querystr, ok := mapreq["query"].(string)
	if !ok {
		return MakeGraphQLErrorResult(pb.CODE_HTTP_NO_QUERY)
	}

	mapval, ok1 := mapreq["variables"].(map[string]interface{})
	if !ok1 {
		mapval = nil
		// return MakeGraphQLErrorResult(pb.CODE_HTTP_VARIABLE_ERR)
	}

	// curdb := s.anka.MgrDB.GetDB(ankadbname)
	curctx := context.WithValue(r.Context(), interface{}("ankadb"), s.anka)

	result1, err := s.anka.logic.OnQuery(curctx, querystr, mapval)
	if err != nil {
		return MakeGraphQLErrorResult(ankadberr.GetErrCode(err))
	}

	return result1
}

func (s *ankaHTTPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "/graphql" {
		result := s.procGraphQL(w, r)

		json.NewEncoder(w).Encode(result)
	}
}

// newHTTPServer -
func newHTTPServer(anka *AnkaDB) (*ankaHTTPServer, error) {
	lis, err := net.Listen("tcp", anka.cfg.AddrHTTP)
	if err != nil {
		return nil, err
	}

	// http.Serve(lis)

	s := &ankaHTTPServer{
		anka:     anka,
		lis:      lis,
		chanServ: make(chan int),
	}

	// pb.RegisterAnkaDBServServer(grpcServ, s)

	return s, nil
}

func (s *ankaHTTPServer) start() (err error) {
	err = http.Serve(s.lis, s)

	s.chanServ <- 0

	return
}

func (s *ankaHTTPServer) stop() {
	s.lis.Close()

	return
}
