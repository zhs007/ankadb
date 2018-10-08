package ankadb

import (
	"github.com/zhs007/ankadb/database"
)

// DBMgr -
type DBMgr interface {
	AddDB(cfg DBConfig) error
	GetDB(name string) database.Database
}

// newDBMgr -
func newDBMgr(lstDB []DBConfig) (DBMgr, error) {
	mgr := &dbMgr{
		mapDB: make(map[string]database.Database),
	}

	for _, val := range lstDB {
		err := mgr.AddDB(val)
		if err != nil {
			return nil, err
		}
	}

	return mgr, nil
}

type dbMgr struct {
	mapDB map[string]database.Database
}

// AddDB -
func (mgr *dbMgr) AddDB(cfg DBConfig) error {
	if cfg.Engine == "leveldb" {
		db, err := database.NewAnkaLDB(cfg.PathDB, 16, 16)
		if err != nil {
			return err
		}

		mgr.mapDB[cfg.Name] = db
	}

	return nil
}

// GetDB -
func (mgr *dbMgr) GetDB(name string) database.Database {
	if db, ok := mgr.mapDB[name]; ok {
		return db
	}

	return nil
}
