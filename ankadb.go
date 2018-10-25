package ankadb

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/graphql-go/graphql"
	"github.com/zhs007/ankadb/err"
	pb "github.com/zhs007/ankadb/proto"
)

// AnkaDB -
type AnkaDB struct {
	MgrDB      DBMgr
	serv       *ankaServer
	servHTTP   *ankaHTTPServer
	cfg        Config
	logic      DBLogic
	chanSignal chan os.Signal
}

// NewAnkaDB -
func NewAnkaDB(cfg Config, logic DBLogic) (*AnkaDB, error) {
	// return nil
	dbmgr, err := NewDBMgr(cfg.ListDB)
	if err != nil {
		return nil, ankadberr.NewError(pb.CODE_INIT_NEW_DBMGR_ERR)
	}

	anka := AnkaDB{
		MgrDB:      dbmgr,
		cfg:        cfg,
		logic:      logic,
		chanSignal: make(chan os.Signal, 1),
	}

	if cfg.AddrGRPC != "" {
		serv, err := newServer(&anka)
		if err != nil {
			return nil, ankadberr.NewError(pb.CODE_INIT_NEW_GRPCSERV_ERR)
		}

		anka.serv = serv
	}

	if cfg.AddrHTTP != "" {
		httpserv, err := newHTTPServer(&anka)
		if err != nil {
			return nil, ankadberr.NewError(pb.CODE_INIT_NEW_HTTPSERV_ERR)
		}

		anka.servHTTP = httpserv
	}

	signal.Notify(anka.chanSignal, os.Interrupt, os.Kill, syscall.SIGSTOP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGTSTP)

	return &anka, nil
}

// Start -
func (anka *AnkaDB) Start() {
	if anka.serv == nil && anka.servHTTP == nil {
		return
	}

	if anka.serv != nil {
		go anka.serv.start()
	}

	if anka.servHTTP != nil {
		go anka.servHTTP.start()
	}

	anka.waitEnd()
}

// Stop -
func (anka *AnkaDB) Stop() {
	anka.serv.stop()
	if anka.servHTTP != nil {
		anka.servHTTP.stop()
	}
}

func (anka *AnkaDB) waitEnd() {
	if anka.servHTTP != nil {
		exitnums := -2
		for {
			select {
			case signal := <-anka.chanSignal:
				fmt.Printf("get signal " + signal.String() + "\n")
				anka.Stop()
			case <-anka.serv.chanServ:
				fmt.Printf("grpcserv exit \n")
				exitnums++
				if exitnums >= 0 {
					return
				}
			case <-anka.servHTTP.chanServ:
				fmt.Printf("httpserv exit \n")
				exitnums++
				if exitnums >= 0 {
					return
				}
			}
		}
	}

	for {
		select {
		case signal := <-anka.chanSignal:
			fmt.Printf("get signal " + signal.String() + "\n")
			anka.Stop()
		case <-anka.serv.chanServ:
			return
		}
	}
}

// LocalQuery - local query
func (anka *AnkaDB) LocalQuery(ctx context.Context, request string, values map[string]interface{}) (*graphql.Result, error) {
	return anka.logic.OnQuery(ctx, request, values)
}
