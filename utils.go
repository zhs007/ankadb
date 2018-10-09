package ankadb

import (
	"context"

	"github.com/zhs007/ankadb/database"
)

// GetContextValueDatabase -
func GetContextValueDatabase(ctx context.Context, key interface{}) database.Database {
	val := ctx.Value(key)
	if val == nil {
		return nil
	}

	if db, ok := val.(database.Database); ok {
		return db
	}

	return nil
}
