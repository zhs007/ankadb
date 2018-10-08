package ankadb

import "github.com/graphql-go/graphql"

// DBLogic -
type DBLogic interface {
	OnQuery(request string, values map[string]interface{}) (*graphql.Result, error)
}
