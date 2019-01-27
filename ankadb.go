package ankadb

import (
	"context"

	"github.com/graphql-go/graphql"
)

// AnkaDB - AnkaDB interface
type AnkaDB interface {
	// Start - start service
	Start(ctx context.Context) error
	// Stop - stop service
	Stop() error

	// LocalQuery - local query
	LocalQuery(ctx context.Context, request string, values map[string]interface{}) (*graphql.Result, error)

	// Get - get value
	Get(ctx context.Context, dbname string, key string) ([]byte, error)
	// Set - set value
	Set(ctx context.Context, dbname string, key string, value []byte) error

	// GetConfig - get config
	GetConfig() *Config
	// GetLogic - get DBLogic
	GetLogic() DBLogic
	// GetDBMgr - get DBMgr
	GetDBMgr() DBMgr
}

// ankaDB - An implementation for AnkaDB
type ankaDB struct {
	mgrDB    DBMgr
	serv     *ankaServer
	servHTTP *ankaHTTPServer
	cfg      Config
	logic    DBLogic
}

// NewAnkaDB -
func NewAnkaDB(cfg Config, logic DBLogic) (AnkaDB, error) {
	// return nil
	dbmgr, err := NewDBMgr(cfg.ListDB)
	if err != nil {
		return nil, err
	}

	anka := &ankaDB{
		mgrDB: dbmgr,
		cfg:   cfg,
		logic: logic,
	}

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
func (anka *ankaDB) Stop() error {
	if anka.serv != nil {
		anka.serv.stop()
	}

	if anka.servHTTP != nil {
		anka.servHTTP.stop()
	}

	return nil
}

// LocalQuery - local query
func (anka *ankaDB) LocalQuery(ctx context.Context, request string, values map[string]interface{}) (*graphql.Result, error) {
	curctx := context.WithValue(ctx, interface{}("ankadb"), anka)

	return anka.logic.OnQuery(curctx, request, values)
}

// GetConfig - get config
func (anka *ankaDB) GetConfig() *Config {
	return &anka.cfg
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

	return db.Get([]byte(key))
}

// Set - set value
func (anka *ankaDB) Set(ctx context.Context, dbname string, key string, value []byte) error {
	db := anka.mgrDB.GetDB(dbname)
	if db == nil {
		return ErrNotFoundDB
	}

	return db.Put([]byte(key), []byte(value))
}
