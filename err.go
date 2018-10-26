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
)
