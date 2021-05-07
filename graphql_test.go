package ankadb

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/graphql-go/graphql"
	"github.com/zhs007/ankadb/graphqlext"
	testpb "github.com/zhs007/ankadb/test"
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
					anka := GetContextValueAnkaDB(params.Context, RequestIDKey)
					if anka == nil {
						return nil, ErrCtxAnkaDB
					}

					curdb := anka.GetDatabase("msg")
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
					anka := GetContextValueAnkaDB(params.Context, RequestIDKey)
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
					anka := GetContextValueAnkaDB(params.Context, RequestIDKey)
					if anka == nil {
						return nil, ErrCtxAnkaDB
					}

					curdb := anka.GetDatabase("user")
					if curdb == nil {
						return nil, ErrCtxCurDB
					}

					lstUser := &testpb.UserList{}
					it := curdb.NewIteratorWithPrefix([]byte(prefixKeyUser))
					if it.Error() != nil {
						return nil, it.Error()
					}

					for {
						if it.Valid() {
							cu := &testpb.User{}
							err := proto.Unmarshal(it.Value(), cu)
							if err != nil {
								return nil, err
							}

							// fmt.Printf("key-%v value-%v\n", it.Key(), cu)

							lstUser.Users = append(lstUser.Users, cu)
						}

						if !it.Next() {
							break
						}
					}

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
					anka := GetContextValueAnkaDB(params.Context, RequestIDKey)
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
				anka := GetContextValueAnkaDB(params.Context, RequestIDKey)
				if anka == nil {
					return nil, ErrCtxAnkaDB
				}

				curdb := anka.GetDatabase("msg")
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
				anka := GetContextValueAnkaDB(params.Context, RequestIDKey)
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

// // resultUpdUser - updUser
// type resultUpdUser struct {
// 	UpdUser struct {
// 		UserID string `json:"userID"`
// 	} `json:"updUser"`
// }

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

const queryUserWithUserName = `query UserWithUserName($userName: ID!) {
	userWithUserName(userName: $userName) {
		nickName
		userID
		userName
	}
}`

// GetUser - get user
func (db *testDB) GetUserWithUserName(userName string) (*testpb.User, error) {
	if db.db == nil {
		return nil, ErrNotInit
	}

	params := make(map[string]interface{})
	params["userName"] = userName

	result, err := db.db.Query(context.Background(), queryUserWithUserName, params)
	if err != nil {
		return nil, err
	}

	err = GetResultError(result)
	if err != nil {
		return nil, err
	}

	// fmt.Printf("%v", result)

	user := &testpb.User{}
	err = MakeMsgFromResultEx(result, "userWithUserName", user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

const queryUsers = `{
	users {
		users {
			nickName
			userID
			userName
		}
	}
}`

// GetUsers - get users
func (db *testDB) GetUsers() (*testpb.UserList, error) {
	if db.db == nil {
		return nil, ErrNotInit
	}

	// params := make(map[string]interface{})
	// params["userID"] = userID

	result, err := db.db.Query(context.Background(), queryUsers, nil)
	if err != nil {
		return nil, err
	}

	err = GetResultError(result)
	if err != nil {
		return nil, err
	}

	// fmt.Printf("%v", result)

	users := &testpb.UserList{}
	err = MakeMsgFromResultEx(result, "users", users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

const queryUpdMsg = `mutation UpdMsg($msg: MessageInput!) {
	updMsg(msg: $msg) {
		chatID
	}
}`

// UpdMsg - update msg
func (db *testDB) UpdMsg(msg *testpb.Message) (string, error) {
	if db.db == nil {
		return "", ErrNotInit
	}

	params := make(map[string]interface{})
	params["msg"] = Msg2Map(msg)

	result, err := db.db.Query(context.Background(), queryUpdMsg, params)
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
	retmsg := &testpb.Message{}
	err = MakeMsgFromResultEx(result, "updMsg", retmsg)
	if err != nil {
		return "", err
	}

	return retmsg.ChatID, nil
}

const queryMsg = `query Msg($chatID: ID!) {
	msg(chatID: $chatID) {
		chatID
		from{
			userID
		}
		to{
			userID
		}
		text
		timeStamp
	}
}`

// GetMsg - get message
func (db *testDB) GetMsg(chatID string) (*testpb.Message, error) {
	if db.db == nil {
		return nil, ErrNotInit
	}

	params := make(map[string]interface{})
	params["chatID"] = chatID

	result, err := db.db.Query(context.Background(), queryMsg, params)
	if err != nil {
		return nil, err
	}

	err = GetResultError(result)
	if err != nil {
		return nil, err
	}

	// fmt.Printf("%v", result)

	msg := &testpb.Message{}
	err = MakeMsgFromResultEx(result, "msg", msg)
	if err != nil {
		return nil, err
	}

	return msg, nil
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

		user1, err := tdb.GetUserWithUserName(username)
		if err != nil {
			t.Fatalf("Test_GraphQL GetUserWithUserName err %v", err)

			return
		}

		if user1.UserID != userid {
			t.Fatalf("Test_GraphQL GetUserWithUserName UserID err %v", user1.UserID)

			return
		}

		if user1.NickName != nickname {
			t.Fatalf("Test_GraphQL GetUserWithUserName NickName err %v", user1.NickName)

			return
		}

		if user1.UserName != username {
			t.Fatalf("Test_GraphQL GetUserWithUserName UserName err %v", user1.UserName)

			return
		}
	}

	users, err := tdb.GetUsers()
	if err != nil {
		t.Fatalf("Test_GraphQL GetUsers err %v", err)
	}

	if len(users.Users) != 100 {
		t.Fatalf("Test_GraphQL GetUsers len err %v", len(users.Users))
	}

	sort.Slice(users.Users, func(i, j int) bool {
		iuid, _ := strconv.Atoi(users.Users[i].UserID)
		juid, _ := strconv.Atoi(users.Users[j].UserID)
		return iuid < juid
	})

	for i := 0; i < 100; i++ {
		nickname := fmt.Sprintf("user %d", i)
		userid := fmt.Sprintf("%d", (i + 1))
		username := fmt.Sprintf("user%d", i)

		if users.Users[i].UserID != userid {
			t.Fatalf("Test_GraphQL GetUsers UserID err %v", users.Users[i].UserID)

			return
		}

		if users.Users[i].NickName != nickname {
			t.Fatalf("Test_GraphQL GetUsers NickName err %v", users.Users[i].NickName)

			return
		}

		if users.Users[i].UserName != username {
			t.Fatalf("Test_GraphQL GetUsers UserName err %v", users.Users[i].UserName)

			return
		}
	}

	chatid, err := tdb.UpdMsg(&testpb.Message{
		ChatID: "00001",
		From: &testpb.User{
			UserID: "1",
		},
		To: &testpb.User{
			UserID: "2",
		},
		Text:      "text",
		TimeStamp: 1234567,
	})
	if err != nil {
		t.Fatalf("Test_GraphQL UpdMsg err %v", err)
	}

	if chatid != "00001" {
		t.Fatalf("Test_GraphQL UpdMsg chatid err %v", chatid)

		return
	}

	msg, err := tdb.GetMsg("00001")
	if err != nil {
		t.Fatalf("Test_GraphQL GetMsg err %v", err)
	}

	if msg.Text != "text" {
		t.Fatalf("Test_GraphQL GetMsg msg.Text err %v", msg.Text)

		return
	}

	t.Logf("Test_GraphQL OK")
}
