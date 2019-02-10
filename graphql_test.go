package ankadb

import (
	"context"
	"fmt"
	"testing"

	"github.com/graphql-go/graphql"
	"github.com/zhs007/ankadb/graphqlext"
	"github.com/zhs007/ankadb/test"
)

const prefixKeyMessage = "msg:"
const prefixKeyUser = "user:"
const prefixKeyUserName = "uname:"

func makeMessageKey(chatID string) string {
	return prefixKeyMessage + chatID
}

func makeUserKey(userID string) string {
	return prefixKeyUser + userID
}

func makeUserNameKey(userName string) string {
	return prefixKeyUserName + userName
}

// inputTypeMessage - Message
//		you can see graphql_test.graphql
var inputTypeMessage = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "MessageInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"chatID": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
			"from": &graphql.InputObjectFieldConfig{
				Type: typeUser,
			},
			"to": &graphql.InputObjectFieldConfig{
				Type: typeUser,
			},
			"text": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"timeStamp": &graphql.InputObjectFieldConfig{
				Type: graphqlext.Int64,
			},
		},
	},
)

// inputTypeUser - User
//		you can see graphql_test.graphql
var inputTypeUser = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "UserInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"nickName": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"userID": &graphql.InputObjectFieldConfig{
				Type: graphql.ID,
			},
			"userName": &graphql.InputObjectFieldConfig{
				Type: graphql.ID,
			},
		},
	},
)

// typeUser - User
//		you can see graphql_test.graphql
var typeUser = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"nickName": &graphql.Field{
				Type: graphql.String,
			},
			"userID": &graphql.Field{
				Type: graphql.ID,
			},
			"userName": &graphql.Field{
				Type: graphql.ID,
			},
		},
	},
)

// typeMessage - Message
//		you can see graphql_test.graphql
var typeMessage = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Message",
		Fields: graphql.Fields{
			"chatID": &graphql.Field{
				Type: graphql.NewNonNull(graphql.ID),
			},
			"from": &graphql.Field{
				Type: typeUser,
			},
			"to": &graphql.Field{
				Type: typeUser,
			},
			"text": &graphql.Field{
				Type: graphql.String,
			},
			"timeStamp": &graphql.Field{
				Type: graphqlext.Int64,
			},
		},
	},
)

// typeUserList - UserList
//		you can see graphql_test.graphql
var typeUserList = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "UserList",
		Fields: graphql.Fields{
			"users": &graphql.Field{
				Type: graphql.NewList(typeUser),
			},
		},
	},
)

// typeQuery - Query
//		you can see graphql_test.graphql
var typeQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"msg": &graphql.Field{
				Type: typeMessage,
				Args: graphql.FieldConfigArgument{
					"chatID": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.ID),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					anka := GetContextValueAnkaDB(params.Context, interface{}("ankadb"))
					if anka == nil {
						return nil, ErrCtxAnkaDB
					}

					curdb := anka.GetDatabase("user")
					if curdb == nil {
						return nil, ErrCtxCurDB
					}

					chatID := params.Args["chatID"].(string)

					msg := &testpb.Message{}
					err := GetMsgFromDB(curdb, []byte(makeMessageKey(chatID)), msg)
					if err != nil {
						return nil, err
					}

					return msg, nil
				},
			},
			"user": &graphql.Field{
				Type: typeUser,
				Args: graphql.FieldConfigArgument{
					"userID": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.ID),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					anka := GetContextValueAnkaDB(params.Context, interface{}("ankadb"))
					if anka == nil {
						return nil, ErrCtxAnkaDB
					}

					curdb := anka.GetDatabase("user")
					if curdb == nil {
						return nil, ErrCtxCurDB
					}

					userID := params.Args["userID"].(string)

					user := &testpb.User{}
					err := GetMsgFromDB(curdb, []byte(makeUserKey(userID)), user)
					if err != nil {
						return nil, err
					}

					return user, nil
				},
			},
			"users": &graphql.Field{
				Type: typeUserList,
				Args: graphql.FieldConfigArgument{},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					anka := GetContextValueAnkaDB(params.Context, interface{}("ankadb"))
					if anka == nil {
						return nil, ErrCtxAnkaDB
					}

					curdb := anka.GetDatabase("user")
					if curdb == nil {
						return nil, ErrCtxCurDB
					}

					lstUser := &testpb.UserList{}

					return lstUser, nil
				},
			},
			"userWithUserName": &graphql.Field{
				Type: typeUser,
				Args: graphql.FieldConfigArgument{
					"userName": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.ID),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					anka := GetContextValueAnkaDB(params.Context, interface{}("ankadb"))
					if anka == nil {
						return nil, ErrCtxAnkaDB
					}

					curdb := anka.GetDatabase("user")
					if curdb == nil {
						return nil, ErrCtxCurDB
					}

					userName := params.Args["userName"].(string)
					uid, err := curdb.Get([]byte(makeUserNameKey(userName)))
					if err != nil {
						return nil, err
					}

					user := &testpb.User{}
					err = GetMsgFromDB(curdb, []byte(makeUserKey(string(uid))), user)
					if err != nil {
						return nil, err
					}

					return user, nil
				},
			},
		},
	},
)

// typeMutation - Mutation
//		you can see graphql_test.graphql
var typeMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"updMsg": &graphql.Field{
			Type:        typeMessage,
			Description: "update message",
			Args: graphql.FieldConfigArgument{
				"msg": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(inputTypeMessage),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				anka := GetContextValueAnkaDB(params.Context, interface{}("ankadb"))
				if anka == nil {
					return nil, ErrCtxAnkaDB
				}

				curdb := anka.GetDatabase("user")
				if curdb == nil {
					return nil, ErrCtxCurDB
				}

				msg := &testpb.Message{}
				err := GetMsgFromParam(params, "msg", msg)
				if err != nil {
					return nil, err
				}

				err = PutMsg2DB(curdb, []byte(makeMessageKey(msg.ChatID)), msg)
				if err != nil {
					return nil, err
				}

				return msg, nil
			},
		},
		"updUser": &graphql.Field{
			Type:        typeUser,
			Description: "update user",
			Args: graphql.FieldConfigArgument{
				"user": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(inputTypeUser),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				anka := GetContextValueAnkaDB(params.Context, interface{}("ankadb"))
				if anka == nil {
					return nil, ErrCtxAnkaDB
				}

				curdb := anka.GetDatabase("user")
				if curdb == nil {
					return nil, ErrCtxCurDB
				}

				user := &testpb.User{}
				err := GetMsgFromParam(params, "user", user)
				if err != nil {
					return nil, err
				}

				err = PutMsg2DB(curdb, []byte(makeUserKey(user.UserID)), user)
				if err != nil {
					return nil, err
				}

				if user.UserName != "" {
					err = curdb.Put([]byte(makeUserNameKey(user.UserName)), []byte(user.UserID))
					if err != nil {
						return nil, err
					}
				}

				return user, nil
			},
		},
	},
})

// testDB - testdb
type testDB struct {
	db AnkaDB
}

// newTestDB - new testDB
func newTestDB(cfg *Config) (*testDB, error) {
	basedblogic, err := NewBaseDBLogic(graphql.SchemaConfig{
		Query:    typeQuery,
		Mutation: typeMutation,
	})
	if err != nil {
		return nil, err
	}

	db, err := NewAnkaDB(cfg, basedblogic)
	if err != nil {
		return nil, err
	}

	return &testDB{
		db: db,
	}, nil
}

const queryUpdUser = `mutation UpdUser($user: UserInput!) {
	updUser(user: $user) {
		userID
	}
}`

// resultUpdUser - updUser
type resultUpdUser struct {
	UpdUser struct {
		UserID string `json:"userID"`
	} `json:"updUser"`
}

// UpdUser - update user
func (db *testDB) UpdUser(user *testpb.User) (string, error) {
	if db.db == nil {
		return "", ErrNotInit
	}

	params := make(map[string]interface{})
	params["user"] = Msg2Map(user)

	result, err := db.db.Query(context.Background(), queryUpdUser, params)
	if err != nil {
		return "", err
	}

	err = GetResultError(result)
	if err != nil {
		return "", err
	}

	// fmt.Printf("%v", result)

	// uu := &resultUpdUser{}
	// err = MakeObjFromResult(result, uu)
	// if err != nil {
	// 	return "", err
	// }
	retuser := &testpb.User{}
	err = MakeMsgFromResultEx(result, "updUser", retuser)
	if err != nil {
		return "", err
	}

	return retuser.UserID, nil
}

const queryUser = `query User($userID: ID!) {
	user(userID: $userID) {
		nickName
		userID
		userName
	}
}`

// resultUser - user
type resultUser struct {
	User struct {
		NickName string `json:"nickName"`
		UserID   string `json:"userID"`
		UserName string `json:"userName"`
	} `json:"user"`
}

// GetUser - get user
func (db *testDB) GetUser(userID string) (*testpb.User, error) {
	if db.db == nil {
		return nil, ErrNotInit
	}

	params := make(map[string]interface{})
	params["userID"] = userID

	result, err := db.db.Query(context.Background(), queryUser, params)
	if err != nil {
		return nil, err
	}

	err = GetResultError(result)
	if err != nil {
		return nil, err
	}

	// fmt.Printf("%v", result)

	user := &testpb.User{}
	err = MakeMsgFromResultEx(result, "user", user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func Test_GraphQL(t *testing.T) {
	cfg, err := LoadConfig("./test/graphql.yaml")
	if err != nil {
		t.Fatalf("Test_GraphQL LoadConfig err %v", err)

		return
	}

	tdb, err := newTestDB(cfg)
	if err != nil {
		t.Fatalf("Test_GraphQL newTestDB err %v", err)

		return
	}

	for i := 0; i < 100; i++ {
		nickname := fmt.Sprintf("user %d", i)
		userid := fmt.Sprintf("%d", (i + 1))
		username := fmt.Sprintf("user%d", i)

		uid, err := tdb.UpdUser(&testpb.User{
			NickName: nickname,
			UserID:   userid,
			UserName: username,
		})
		if err != nil {
			t.Fatalf("Test_GraphQL UpdUser err %v", err)

			return
		}

		if uid != userid {
			t.Fatalf("Test_GraphQL UpdUser uid err %v", uid)

			return
		}
	}

	for i := 0; i < 100; i++ {
		nickname := fmt.Sprintf("user %d", i)
		userid := fmt.Sprintf("%d", (i + 1))
		username := fmt.Sprintf("user%d", i)

		user, err := tdb.GetUser(userid)
		if err != nil {
			t.Fatalf("Test_GraphQL GetUser err %v", err)

			return
		}

		if user.UserID != userid {
			t.Fatalf("Test_GraphQL GetUser UserID err %v", user.UserID)

			return
		}

		if user.NickName != nickname {
			t.Fatalf("Test_GraphQL GetUser NickName err %v", user.NickName)

			return
		}

		if user.UserName != username {
			t.Fatalf("Test_GraphQL GetUser UserName err %v", user.UserName)

			return
		}
	}

	t.Logf("Test_GraphQL OK")
}
