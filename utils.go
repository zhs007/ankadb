package ankadb

import (
	"context"
	"encoding/json"

	"github.com/goinggo/mapstructure"
	"github.com/golang/protobuf/proto"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/zhs007/ankadb/database"
)

// GetContextValueAnkaDB -
func GetContextValueAnkaDB(ctx context.Context, key interface{}) *AnkaDB {
	val := ctx.Value(key)
	if val == nil {
		return nil
	}

	if db, ok := val.(*AnkaDB); ok {
		return db
	}

	return nil
}

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
func MakeGraphQLErrorResult(err error) *graphql.Result {
	result := graphql.Result{}

	gqlerr := gqlerrors.FormattedError{
		Message: err.Error(),
	}

	result.Errors = append(result.Errors, gqlerr)

	return &result
}

// PutMsg2DB - put protobuf message to database
func PutMsg2DB(db database.Database, key []byte, msg proto.Message) error {
	data, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	err = db.Put(key, data)
	if err != nil {
		return err
	}

	return nil
}

// GetMsgFromDB - get protobuf message from database
func GetMsgFromDB(db database.Database, key []byte, msg proto.Message) error {
	buf, err := db.Get(key)
	if err != nil {
		return err
	}

	err = proto.Unmarshal(buf, msg)
	if err != nil {
		return err
	}

	return nil
}

// GetMsgFromParam - get protobuf message from param
func GetMsgFromParam(params graphql.ResolveParams, paramName string, msg proto.Message) error {
	ci := params.Args[paramName].(map[string]interface{})

	if err := mapstructure.Decode(ci, msg); err != nil {
		return err
	}

	return nil
}

// MakeParamsFromMsg - change protobuf message to param
func MakeParamsFromMsg(params map[string]interface{}, paramName string, msg proto.Message) error {
	cv := make(map[string]interface{})

	inrec, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	json.Unmarshal(inrec, &cv)

	params[paramName] = cv

	return nil
}

// MakeObjFromResult - make object from graphql.Result
func MakeObjFromResult(result *graphql.Result, obj interface{}) error {
	if err := mapstructure.Decode(result.Data, obj); err != nil {
		return err
	}

	return nil
}

// MakeMsgFromResult - make protobuf object from graphql.Result
func MakeMsgFromResult(result *graphql.Result, msg proto.Message) error {
	if err := mapstructure.Decode(result.Data, msg); err != nil {
		return err
	}

	return nil
}
