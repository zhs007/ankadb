package ankadb

import (
	"context"

	"github.com/graphql-go/graphql"
	"github.com/zhs007/ankadb/database"
)

// FuncAnkaDBForEach - func event
type FuncAnkaDBForEach func(key string, value []byte) error

// AnkaDB - AnkaDB interface
type AnkaDB interface {
	// Start - start service
	Start(ctx context.Context) error
	// // Stop - stop service
	// Stop() error

	// Query - query
	Query(ctx context.Context, request string, values map[string]interface{}) (*graphql.Result, error)

	// SetQueryTemplate - set query template
	SetQueryTemplate(templateName string, request string) error

	// Get - get value
	Get(ctx context.Context, dbname string, key string) ([]byte, error)
	// Set - set value
	Set(ctx context.Context, dbname string, key string, value []byte) error
	// Delete - deletes the key from database
	Delete(ctx context.Context, dbname string, key string) error

	// ForEachWithPrefix - for each with prefix
	ForEachWithPrefix(ctx context.Context, dbname string, prefix string, foreach FuncAnkaDBForEach) error

	// RegEventFunc - register a function for event
	RegEventFunc(event string, eventfunc FuncAnkaDBEvent) error

	// GetConfig - get config
	GetConfig() *Config
	// GetLogic - get DBLogic
	GetLogic() DBLogic
	// GetDBMgr - get DBMgr
	GetDBMgr() DBMgr

	// GetDatabase - get databse with dbname
	GetDatabase(dbname string) database.Database
	// AddDatabase - add database
	AddDatabase(dbcfg DBConfig) error
}

// ankaDB - An implementation for AnkaDB
type ankaDB struct {
	mgrDB    DBMgr
	serv     *ankaServer
	servHTTP *ankaHTTPServer
	cfg      *Config
	logic    DBLogic
	mgrEvent *eventMgr
}

// NewAnkaDB -
func NewAnkaDB(cfg *Config, logic DBLogic) (AnkaDB, error) {
	// return nil
	dbmgr, err := NewDBMgr(cfg.PathDBRoot, cfg.ListDB)
	if err != nil {
		return nil, err
	}

	anka := &ankaDB{
		mgrDB: dbmgr,
		cfg:   cfg,
		logic: logic,
	}

	anka.mgrEvent = newEventMgr(anka)

	return anka, nil
}

// Start -
func (anka *ankaDB) Start(ctx context.Context) error {
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
		grpcctx, grpccancel = context.WithCancel(ctx)

		go anka.serv.start(grpcctx)
	}

	var httpctx context.Context
	var httpcancel context.CancelFunc

	if anka.servHTTP != nil {
		httpctx, httpcancel = context.WithCancel(ctx)

		go anka.servHTTP.start(httpctx)
	}

	anka.mgrEvent.onAnkaDBEvent(ctx, EventOnStarted)

	select {
	case <-ctx.Done():
		if grpccancel != nil {
			grpccancel()
		}

		if httpcancel != nil {
			httpcancel()
		}

		anka.stop()

		return nil
	}
}

// stop -
func (anka *ankaDB) stop() error {
	if anka.serv != nil {
		anka.serv.stop()
	}

	if anka.servHTTP != nil {
		anka.servHTTP.stop()
	}

	return nil
}

// Query - query
func (anka *ankaDB) Query(ctx context.Context, request string, values map[string]interface{}) (*graphql.Result, error) {
	curctx := context.WithValue(ctx, interface{}("ankadb"), anka)

	return anka.logic.Query(curctx, request, values)
}

// SetQueryTemplate - set query template
func (anka *ankaDB) SetQueryTemplate(templateName string, request string) error {
	err := anka.logic.SetQueryTemplate(templateName, request)
	if err != nil {
		return err
	}

	return nil
}

// GetConfig - get config
func (anka *ankaDB) GetConfig() *Config {
	return anka.cfg
}

// GetLogic - get DBLogic
func (anka *ankaDB) GetLogic() DBLogic {
	return anka.logic
}

// GetDBMgr - get DBMgr
func (anka *ankaDB) GetDBMgr() DBMgr {
	return anka.mgrDB
}

// Get - get value
func (anka *ankaDB) Get(ctx context.Context, dbname string, key string) ([]byte, error) {
	db := anka.mgrDB.GetDB(dbname)
	if db == nil {
		return nil, ErrNotFoundDB
	}

	val, err := db.Get([]byte(key))
	if err == database.ErrNotFound {
		return nil, ErrNotFoundKey
	}

	return val, err
}

// Set - set value
func (anka *ankaDB) Set(ctx context.Context, dbname string, key string, value []byte) error {
	db := anka.mgrDB.GetDB(dbname)
	if db == nil {
		return ErrNotFoundDB
	}

	return db.Put([]byte(key), []byte(value))
}

// Delete - deletes the key from database
func (anka *ankaDB) Delete(ctx context.Context, dbname string, key string) error {
	db := anka.mgrDB.GetDB(dbname)
	if db == nil {
		return ErrNotFoundDB
	}

	return db.Delete([]byte(key))
}

// ForEachWithPrefix - for each with prefix
func (anka *ankaDB) ForEachWithPrefix(ctx context.Context, dbname string, prefix string, foreach FuncAnkaDBForEach) error {
	db := anka.mgrDB.GetDB(dbname)
	if db == nil {
		return ErrNotFoundDB
	}

	it := db.NewIteratorWithPrefix([]byte(prefix))
	if it.Error() != nil {
		return it.Error()
	}

	for {
		if it.Valid() {
			err := foreach(string(it.Key()), it.Value())
			if err != nil {
				it.Release()

				return err
			}
		}

		if !it.Next() {
			break
		}
	}

	it.Release()

	return nil
}

// RegEventFunc - register a function for event
func (anka *ankaDB) RegEventFunc(event string, eventfunc FuncAnkaDBEvent) error {
	return anka.mgrEvent.regAnkaDBEventFunc(event, eventfunc)
}

// GetDatabase - get databse with dbname
func (anka *ankaDB) GetDatabase(dbname string) database.Database {
	return anka.mgrDB.GetDB(dbname)
}

// AddDatabase - add a databse
func (anka *ankaDB) AddDatabase(dbcfg DBConfig) error {
	return anka.mgrDB.AddDB(dbcfg)
}
