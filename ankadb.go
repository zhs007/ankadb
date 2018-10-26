package ankadb

import (
	"context"
	"os"

	"github.com/graphql-go/graphql"
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
		return nil, err
	}

	anka := AnkaDB{
		MgrDB:      dbmgr,
		cfg:        cfg,
		logic:      logic,
		chanSignal: make(chan os.Signal, 1),
	}

	// if cfg.AddrGRPC != "" {
	// 	serv, err := newServer(&anka)
	// 	if err != nil {
	// 		return nil, ankadberr.NewError(pb.CODE_INIT_NEW_GRPCSERV_ERR)
	// 	}

	// 	anka.serv = serv
	// }

	// if cfg.AddrHTTP != "" {
	// 	httpserv, err := newHTTPServer(&anka)
	// 	if err != nil {
	// 		return nil, ankadberr.NewError(pb.CODE_INIT_NEW_HTTPSERV_ERR)
	// 	}

	// 	anka.servHTTP = httpserv
	// }

	// signal.Notify(anka.chanSignal, os.Interrupt, os.Kill, syscall.SIGSTOP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGTSTP)

	return &anka, nil
}

// Start -
func (anka *AnkaDB) Start(ctx context.Context) error {
	if anka.cfg.AddrGRPC != "" {
		serv, err := newServer(anka)
		if err != nil {
			return err
		}

		anka.serv = serv
	}

	if anka.cfg.AddrHTTP != "" {
		httpserv, err := newHTTPServer(anka)
		if err != nil {
			return err
		}

		anka.servHTTP = httpserv
	}

	if anka.serv == nil && anka.servHTTP == nil {
		return nil
	}

	var grpcctx context.Context
	var grpccancel context.CancelFunc

	if anka.serv != nil {
		grpcctx, grpccancel = context.WithCancel(context.Background())

		go anka.serv.start(grpcctx)
	}

	var httpctx context.Context
	var httpcancel context.CancelFunc

	if anka.servHTTP != nil {
		httpctx, httpcancel = context.WithCancel(context.Background())

		go anka.servHTTP.start(httpctx)
	}

	select {
	case <-ctx.Done():
		if grpccancel != nil {
			grpccancel()
		}

		if httpcancel != nil {
			httpcancel()
		}

		anka.Stop()

		return nil
	}
}

// Stop -
func (anka *AnkaDB) Stop() {
	if anka.serv != nil {
		anka.serv.stop()
	}

	if anka.servHTTP != nil {
		anka.servHTTP.stop()
	}
}

// func (anka *AnkaDB) waitEnd() {
// 	if anka.servHTTP != nil {
// 		exitnums := -2
// 		for {
// 			select {
// 			case signal := <-anka.chanSignal:
// 				fmt.Printf("get signal " + signal.String() + "\n")
// 				anka.Stop()
// 			case <-anka.serv.chanServ:
// 				fmt.Printf("grpcserv exit \n")
// 				exitnums++
// 				if exitnums >= 0 {
// 					return
// 				}
// 			case <-anka.servHTTP.chanServ:
// 				fmt.Printf("httpserv exit \n")
// 				exitnums++
// 				if exitnums >= 0 {
// 					return
// 				}
// 			}
// 		}
// 	}

// 	for {
// 		select {
// 		case signal := <-anka.chanSignal:
// 			fmt.Printf("get signal " + signal.String() + "\n")
// 			anka.Stop()
// 		case <-anka.serv.chanServ:
// 			return
// 		}
// 	}
// }

// LocalQuery - local query
func (anka *AnkaDB) LocalQuery(ctx context.Context, request string, values map[string]interface{}) (*graphql.Result, error) {
	return anka.logic.OnQuery(ctx, request, values)
}
