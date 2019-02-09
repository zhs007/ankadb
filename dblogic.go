package ankadb

import (
	"context"

	"github.com/graphql-go/graphql"
	pb "github.com/zhs007/ankadb/proto"
)

// FuncOnQueryStream - use in DBLogic.OnQueryStream
type FuncOnQueryStream func(*pb.ReplyQuery)

// DBLogic -
type DBLogic interface {
	// GetScheme - get GraphQL scheme
	GetScheme() *graphql.Schema

	// OnQuery - query graph request string
	OnQuery(ctx context.Context, request string, values map[string]interface{}) (*graphql.Result, error)
}

// BaseDBLogic - base DBLogic
type BaseDBLogic struct {
	schema *graphql.Schema
}

// NewBaseDBLogic - new BaseDBLogic
func NewBaseDBLogic(cfg graphql.SchemaConfig) (*BaseDBLogic, error) {
	schema, err := graphql.NewSchema(cfg)
	if err != nil {
		return nil, err
	}

	return &BaseDBLogic{
		schema: &schema,
	}, nil
}

// GetScheme - get GraphQL scheme
func (logic *BaseDBLogic) GetScheme() *graphql.Schema {
	return logic.schema
}

// OnQuery - query graph request string
func (logic *BaseDBLogic) OnQuery(ctx context.Context, request string, values map[string]interface{}) (*graphql.Result, error) {
	return nil, nil
}
