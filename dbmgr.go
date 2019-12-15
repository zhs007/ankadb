package ankadb

import (
	"path"

	"github.com/zhs007/ankadb/database"
)

// DBMgr -
type DBMgr interface {
	AddDB(cfg DBConfig) error
	GetDB(name string) database.Database
	GetMgrSnapshot(name string) *SnapshotMgr
}

// NewDBMgr - new DBMgr
func NewDBMgr(rootpath string, lstDB []DBConfig) (DBMgr, error) {
	mgr := &dbMgr{
		mapDB:    make(map[string]*dbObj),
		rootPath: rootpath,
	}

	for _, val := range lstDB {
		err := mgr.AddDB(val)
		if err != nil {
			return nil, err
		}
	}

	return mgr, nil
}

type dbObj struct {
	db          database.Database
	mgrSnapshot *SnapshotMgr
}

type dbMgr struct {
	mapDB    map[string]*dbObj
	rootPath string
}

// AddDB -
func (mgr *dbMgr) AddDB(cfg DBConfig) error {
	_, isok := mgr.mapDB[cfg.Name]
	if isok {
		return ErrDuplicateDB
	}

	if cfg.Engine == "leveldb" {
		db, err := database.NewAnkaLDB(path.Join(mgr.rootPath, cfg.PathDB), 16, 16)
		if err != nil {
			return err
		}

		co := &dbObj{
			db:          db,
			mgrSnapshot: newSnapshotMgr(db),
		}

		co.mgrSnapshot.init()

		mgr.mapDB[cfg.Name] = co
	}

	return nil
}

// GetDB -
func (mgr *dbMgr) GetDB(name string) database.Database {
	if db, ok := mgr.mapDB[name]; ok {
		return db.db
	}

	return nil
}

func (mgr *dbMgr) GetMgrSnapshot(name string) *SnapshotMgr {
	if db, ok := mgr.mapDB[name]; ok {
		return db.mgrSnapshot
	}

	return nil
}
