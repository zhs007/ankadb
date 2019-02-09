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

	t.Logf("Test_Msg2Map OK")
}
