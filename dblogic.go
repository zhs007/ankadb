package ankadb

import (
	"context"
	"errors"

	"github.com/graphql-go/graphql"
	pb "github.com/zhs007/ankadb/proto"
)

// FuncOnQueryStream - use in DBLogic.OnQueryStream
type FuncOnQueryStream func(*pb.ReplyQuery)

// DBLogic -
type DBLogic interface {
	// GetScheme - get GraphQL scheme
	GetScheme() *graphql.Schema

	// Query - query graphql request string
	Query(ctx context.Context, request string, values map[string]interface{}) (*graphql.Result, error)
	// QueryTemplate - query graphql with template
	QueryTemplate(ctx context.Context, templateName string, values map[string]interface{}) (*graphql.Result, error)

	// SetQueryTemplate - set query template
	SetQueryTemplate(templateName string, request string) error
}

// BaseDBLogic - base DBLogic
type BaseDBLogic struct {
	schema       *graphql.Schema
	mgrQueryTemp *queryTemplatesMgr
}

// NewBaseDBLogic - new BaseDBLogic
func NewBaseDBLogic(cfg graphql.SchemaConfig) (*BaseDBLogic, error) {
	schema, err := graphql.NewSchema(cfg)
	if err != nil {
		return nil, err
	}

	return &BaseDBLogic{
		schema:       &schema,
		mgrQueryTemp: newQueryTemplatesMgr(),
	}, nil
}

// GetScheme - get GraphQL scheme
func (logic *BaseDBLogic) GetScheme() *graphql.Schema {
	return logic.schema
}

// Query - query graphql request string
func (logic *BaseDBLogic) Query(ctx context.Context, request string, values map[string]interface{}) (*graphql.Result, error) {
	result := graphql.Do(graphql.Params{
		Schema:         *logic.schema,
		RequestString:  request,
		VariableValues: values,
		Context:        ctx,
	})

	return result, nil
}

// QueryTemplate - query graphql with template
func (logic *BaseDBLogic) QueryTemplate(ctx context.Context, templateName string, values map[string]interface{}) (*graphql.Result, error) {
	return nil, nil
}

// SetQueryTemplate - set query template
func (logic *BaseDBLogic) SetQueryTemplate(templateName string, request string) error {
	err := logic.mgrQueryTemp.setQueryTemplate(logic.schema, templateName, request)
	if err != nil {
		return errors.New(err[0].Error())
	}

	return nil
}
