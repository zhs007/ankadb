package ankadb

import (
	"context"

	"github.com/graphql-go/graphql"
)

// DBLogic -
type DBLogic interface {
	OnQuery(ctx context.Context, request string, values map[string]interface{}) (*graphql.Result, error)
}
