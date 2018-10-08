package ankadb

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// AnkaDB -
type AnkaDB struct {
	dbmgr      DBMgr
	serv       *ankaServer
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
		dbmgr:      dbmgr,
		cfg:        cfg,
		logic:      logic,
		chanSignal: make(chan os.Signal, 1),
	}

	serv, err := newServer(&anka)
	if err != nil {
		return nil
	}

	anka.serv = serv
	signal.Notify(anka.chanSignal, os.Interrupt, os.Kill, syscall.SIGSTOP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGTSTP)

	return &anka
}

// Start -
func (anka *AnkaDB) Start() {
	go anka.serv.start()

	anka.waitEnd()
}

// Stop -
func (anka *AnkaDB) Stop() {
	anka.serv.stop()

}

func (anka *AnkaDB) waitEnd() {
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
