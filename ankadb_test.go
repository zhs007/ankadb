package ankadb

import (
	"context"
	"testing"

	"github.com/graphql-go/graphql"
)

// dblogicTest -
type dblogicTest struct {
}

func (dbl *dblogicTest) OnQuery(ctx context.Context, request string, values map[string]interface{}) (*graphql.Result, error) {
	return nil, nil
}

func (dbl *dblogicTest) OnQueryStream(ctx context.Context, request string, values map[string]interface{}, funcOnQueryStream FuncOnQueryStream) error {
	return nil
}

func Test_AnkaDB(t *testing.T) {
	cfg, err := LoadConfig("./test/test001.yaml")
	if err != nil {
		t.Fatalf("Test_AnkaDB LoadConfig err %v", err)

		return
	}

	anka, err := NewAnkaDB(*cfg, &dblogicTest{})
	if err != nil {
		t.Fatalf("Test_AnkaDB NewAnkaDB err %v", err)

		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	anka.RegEventFunc(EventOnStarted, func(ctx context.Context, anka AnkaDB) error {
		cancel()

		return nil
	})
	// ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// defer cancel()

	err = anka.Start(ctx)
	if err != nil {
		t.Fatalf("Test_AnkaDB Start err %v", err)

		return
	}

	t.Logf("Test_AnkaDB OK")
}
