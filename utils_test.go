package ankadb

import (
	"testing"

	"github.com/zhs007/ankadb/test"
)

func Test_Msg2Map(t *testing.T) {
	user0 := &testpb.User{
		NickName: "user 0",
		UserID:   "1",
		UserName: "user0",
	}

	muser0 := Msg2Map(user0)

	for k, v := range muser0 {
		if k == "nickName" && v == "user 0" {
			continue
		}

		if k == "userID" && v == "1" {
			continue
		}

		if k == "userName" && v == "user0" {
			continue
		}

		t.Fatalf("Test_Msg2Map err k-%v v-%v", k, v)
	}

	// NickName is empty
	user1 := &testpb.User{
		UserID:   "2",
		UserName: "user1",
	}

	muser1 := Msg2Map(user1)

	for k, v := range muser1 {
		// NickName is empty
		if k == "nickName" && v == "" {
			continue
		}

		if k == "userID" && v == "2" {
			continue
		}

		if k == "userName" && v == "user1" {
			continue
		}

		t.Fatalf("Test_Msg2Map err k-%v v-%v", k, v)
	}

	userlist := &testpb.UserList{}
	userlist.Users = append(userlist.Users, user0)
	muserlist := Msg2Map(userlist)

	// fmt.Printf("%v", muserlist)

	for k, v := range muserlist {
		if k == "users" {
			sv := v.([]interface{})
			for i := 0; i < len(sv); i++ {
				for k1, v1 := range sv[i].(map[string]interface{}) {
					if k1 == "nickName" && v1 == "user 0" {
						continue
					}

					if k1 == "userID" && v1 == "1" {
						continue
					}

					if k1 == "userName" && v1 == "user0" {
						continue
					}

					t.Fatalf("Test_Msg2Map err k1-%v v1-%v", k1, v1)
				}
			}

			continue
		}

		t.Fatalf("Test_Msg2Map err k-%v v-%v", k, v)
	}

	t.Logf("Test_Msg2Map OK")
}
