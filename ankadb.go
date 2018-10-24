package ankadb

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
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
func NewAnkaDB(cfg Config, logic DBLogic) *AnkaDB {
	// return nil
	dbmgr, err := NewDBMgr(cfg.ListDB)
	if err != nil {
		return nil
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
			return nil
		}

		anka.serv = serv
	}

	if cfg.AddrHTTP != "" {
		httpserv, err := newHTTPServer(&anka)
		if err != nil {
			return nil
		}

		anka.servHTTP = httpserv
	}

	signal.Notify(anka.chanSignal, os.Interrupt, os.Kill, syscall.SIGSTOP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGTSTP)

	return &anka
}

// Start -
func (anka *AnkaDB) Start() {
	go anka.serv.start()
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
