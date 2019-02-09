package ankadb

import "errors"

var (
	// ErrGRPCListen - grpc service listen err
	ErrGRPCListen = errors.New("grpc service listen err")
	// ErrHTTPListen - http service listen err
	ErrHTTPListen = errors.New("http service listen err")
	// ErrLoadFileReadSize - loadfile invalid file read size
	ErrLoadFileReadSize = errors.New("loadfile invalid file read size")
	// ErrHTTPNoQuery - HTTP no query
	ErrHTTPNoQuery = errors.New("HTTP no query")
	// ErrCtxAnkaDB - context has not ankadb
	ErrCtxAnkaDB = errors.New("context has not ankadb")
	// ErrCtxCurDB - context has not currentdb
	ErrCtxCurDB = errors.New("context has not currentdb")
	// ErrQuertParams - query params err
	ErrQuertParams = errors.New("query params err")
	// ErrQuertResultDecode - query result decode err
	ErrQuertResultDecode = errors.New("query result decode err")
	// ErrCtxSnapshotMgr - query get snapshotMgr from err
	ErrCtxSnapshotMgr = errors.New("query get snapshotMgr from err")
	// ErrNotFoundDB - not found db
	ErrNotFoundDB = errors.New("not found db")
	// ErrNotFoundKey - not found key
	ErrNotFoundKey = errors.New("not found key")
	// ErrInvalidEvent - invalid event
	ErrInvalidEvent = errors.New("invalid event")
	// ErrNotFoundTemplate - not found template
	ErrNotFoundTemplate = errors.New("not found template")
	// ErrNotInit - not initialization
	ErrNotInit = errors.New("not initialization")
)
