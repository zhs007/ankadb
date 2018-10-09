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
	dbmgr, err := newDBMgr(cfg.ListDB)
	if err != nil {
		return nil
	}

	anka := AnkaDB{
		MgrDB:      dbmgr,
		cfg:        cfg,
		logic:      logic,
		chanSignal: make(chan os.Signal, 1),
	}

	serv, err := newServer(&anka)
	if err != nil {
		return nil
	}

	if cfg.AddrHTTP != "" {
		httpserv, err := newHTTPServer(&anka)
		if err != nil {
			return nil
		}

		anka.servHTTP = httpserv
	}

	anka.serv = serv
	signal.Notify(anka.chanSignal, os.Interrupt, os.Kill, syscall.SIGSTOP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGTSTP)

	return &anka
}

// Start -
func (anka *AnkaDB) Start() {
	go anka.serv.start()
	if anka.servHTTP != nil {
		anka.servHTTP.start()
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
				fmt.Printf("get signal " + signal.String())
				anka.Stop()
			case <-anka.serv.chanServ:
				exitnums++
				if exitnums >= 0 {
					return
				}
			case <-anka.servHTTP.chanServ:
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
			fmt.Printf("get signal " + signal.String())
			anka.Stop()
		case <-anka.serv.chanServ:
			return
		}
	}
}

// // NewAnkaDB -
// func NewAnkaDB(cfg DBConfig) (database.Database, error) {
// 	if cfg.Engine == "leveldb" {
// 		return database.NewAnkaLDB(cfg.DBPath, 16, 16)
// 	}

// 	return nil, nil
// }
