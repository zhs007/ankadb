package ankadb

import (
	"github.com/zhs007/ankadb/database"
)

// NewAnkaDB -
func NewAnkaDB(cfg Config) (database.Database, error) {
	if cfg.Engine == "leveldb" {
		return database.NewAnkaLDB(cfg.DBPath, 16, 16)
	}

	return nil, nil
}
