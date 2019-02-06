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
	OnQuery(ctx context.Context, request string, values map[string]interface{}) (*graphql.Result, error)
	OnQueryStream(ctx context.Context, request string, values map[string]interface{}, funcOnQueryStream FuncOnQueryStream) error
}

// BaseDBLogic - base DBLogic
type BaseDBLogic struct {
	schema graphql.Schema
}

// NewBaseDBLogic - new BaseDBLogic
func NewBaseDBLogic(cfg graphql.SchemaConfig) (*BaseDBLogic, error) {
	schema, err := graphql.NewSchema(cfg)
	if err != nil {
		return nil, err
	}

	return &BaseDBLogic{
		schema: schema,
	}, nil
}
