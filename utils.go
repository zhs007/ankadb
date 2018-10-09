package ankadb

import (
	"context"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/zhs007/ankadb/database"
	"github.com/zhs007/ankadb/err"
	pb "github.com/zhs007/ankadb/proto"
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

// MakeGraphQLErrorResult -
func MakeGraphQLErrorResult(code pb.CODE) *graphql.Result {
	result := graphql.Result{}

	err := gqlerrors.FormattedError{
		Message: ankadberr.BuildErrorString(code),
	}

	result.Errors = append(result.Errors, err)

	return &result
}
