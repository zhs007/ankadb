package ankadb

import (
	"github.com/graphql-go/graphql"
	"github.com/zhs007/ankadb/graphqlext"
	"github.com/zhs007/ankadb/test"
)

const prefixKeyMessage = "msg:"
const prefixKeyUser = "user:"
const prefixKeyUserName = "uname:"
const prefixKeyUserScript = "userscript:"
const prefixKeyUserFileTemplate = "filetemplate:"

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
//		you can see dblogic_test.graphql
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
//		you can see dblogic_test.graphql
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
//		you can see dblogic_test.graphql
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
//		you can see dblogic_test.graphql
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
//		you can see dblogic_test.graphql
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

// typeUserList - query
//		you can see dblogic_test.graphql
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

					curdb := anka.GetDatabase("chatbotdb")
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
					// jarvisbase.Debug("query user")

					anka := GetContextValueAnkaDB(params.Context, interface{}("ankadb"))
					if anka == nil {
						return nil, ErrCtxAnkaDB
					}

					curdb := anka.GetDatabase("chatbotdb")
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
					// jarvisbase.Debug("query users")

					anka := GetContextValueAnkaDB(params.Context, interface{}("ankadb"))
					if anka == nil {
						return nil, ErrCtxAnkaDB
					}

					curdb := anka.GetDatabase("chatbotdb")
					if curdb == nil {
						return nil, ErrCtxCurDB
					}

					// mgrSnapshot := anka.MgrDB.GetMgrSnapshot("chatbotdb")
					// if mgrSnapshot == nil {
					// 	return nil, ErrCtxSnapshotMgr
					// }

					// curit := curdb.NewIteratorWithPrefix([]byte(prefixKeyUser))
					// // jarvisbase.Debug("curdb.NewIteratorWithPrefix")
					// for curit.Next() {
					// 	key := curit.Key()
					// 	jarvisbase.Debug("curdb.NewIteratorWithPrefix", zap.String("key", string(key)))
					// }
					// curit.Release()
					// err := curit.Error()
					// if err != nil {
					// 	jarvisbase.Debug("curdb.NewIteratorWithPrefix", zap.Error(err))

					// 	return nil, err
					// }

					// snapshotID := params.Args["snapshotID"].(int64)
					// beginIndex := params.Args["beginIndex"].(int)
					// nums := params.Args["nums"].(int)
					// if beginIndex < 0 || nums <= 0 {
					// 	return nil, ErrQuertParams
					// }

					lstUser := &testpb.UserList{}
					// var pSnapshot *ankadbpb.Snapshot

					// if snapshotID > 0 {
					// 	pSnapshot = mgrSnapshot.Get(snapshotID)
					// } else {
					// 	var err error
					// 	pSnapshot, err = mgrSnapshot.NewSnapshot([]byte(prefixKeyUser))
					// 	if err != nil {
					// 		return nil, ankadb.ErrCtxSnapshotMgr
					// 	}
					// }

					// lstUser.SnapshotID = pSnapshot.SnapshotID
					// lstUser.MaxIndex = int32(len(pSnapshot.Keys))

					// // jarvisbase.Debug("query users", zap.Int32("MaxIndex", lstUser.MaxIndex))

					// curi := beginIndex
					// for ; curi < len(pSnapshot.Keys) && len(lstUser.Users) < nums; curi++ {
					// 	cui := &pb.User{}
					// 	err := ankadb.GetMsgFromDB(curdb, []byte(pSnapshot.Keys[curi]), cui)
					// 	if err == nil {
					// 		// s, err := json.Marshal(cui)
					// 		// if err != nil {
					// 		// 	jarvisbase.Debug("query users", zap.String("user key", pSnapshot.Keys[curi]), zap.Error(err))
					// 		// } else {
					// 		// 	jarvisbase.Debug("query users", zap.String("user key", pSnapshot.Keys[curi]), zap.String("user", string(s)))
					// 		// }

					// 		lstUser.Users = append(lstUser.Users, cui)
					// 	}
					// }

					// lstUser.EndIndex = int32(curi)

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

					curdb := anka.GetDatabase("chatbotdb")
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

				curdb := anka.GetDatabase("chatbotdb")
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

				curdb := anka.GetDatabase("chatbotdb")
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
