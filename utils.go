package ankadb

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/zhs007/ankadb/database"
	"github.com/zhs007/ankadb/err"
	pb "github.com/zhs007/ankadb/proto"
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
func MakeGraphQLErrorResult(code pb.CODE) *graphql.Result {
	result := graphql.Result{}

	err := gqlerrors.FormattedError{
		Message: ankadberr.BuildErrorString(code),
	}

	result.Errors = append(result.Errors, err)

	return &result
}

// PutMsg2DB - put protobuf message to database
func PutMsg2DB(db database.Database, key []byte, msg proto.Message) error {
	data, err := proto.Marshal(msg)
	if err != nil {
		return ankadberr.NewError(pb.CODE_PROTOBUF_ENCODE_ERR)
	}

	err = db.Put(key, data)
	if err != nil {
		return ankadberr.NewError(pb.CODE_DB_PUT_ERR)
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
