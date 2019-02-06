package ankadb

import (
	"sync"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/graphql-go/graphql/language/ast"
)

// queryTemplate - query template
type queryTemplate struct {
	name  string
	query string
	AST   *ast.Document
}

// queryTemplatesMgr - queryTemplates manager
type queryTemplatesMgr struct {
	mapQT sync.Map
}

// newQueryTemplatesMgr - new queryTemplatesMgr
func newQueryTemplatesMgr() *queryTemplatesMgr {
	return &queryTemplatesMgr{
		mapQT: sync.Map{},
	}
}

// getQueryTemplate - get queryTemplate with template name
func (mgr *queryTemplatesMgr) getQueryTemplate(tname string) *queryTemplate {
	v, isok := mgr.mapQT.Load(tname)
	if isok {
		qt, isok := v.(queryTemplate)
		if isok {
			return &qt
		}
	}

	return nil
}

// setQueryTemplate - set queryTemplate
func (mgr *queryTemplatesMgr) setQueryTemplate(schema *graphql.Schema, tname string, query string) []gqlerrors.FormattedError {
	v, isok := mgr.mapQT.Load(tname)
	if isok {
		qt, isok := v.(queryTemplate)
		if isok {
			if qt.query == query {
				return nil
			}

			qt.query = query

			return nil
		}
	}

	qt := queryTemplate{
		name:  tname,
		query: query,
	}

	AST, errs := ParseQuery(schema, query, tname)
	if errs != nil {
		return errs
	}

	qt.AST = AST

	mgr.mapQT.Store(tname, qt)

	return nil
}
