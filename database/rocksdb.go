package database

import (
	"github.com/tecbot/gorocksdb"
)

// ankaRDB - anka RocksDB
type ankaRDB struct {
	dbpath string        // filename for reporting
	db     *gorocksdb.DB // RocksDB instance
}
