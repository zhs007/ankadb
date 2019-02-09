package ankadb

import (
	"context"
	"fmt"
	"testing"
)

func Test_AnkaDB(t *testing.T) {
	cfg, err := LoadConfig("./test/test001.yaml")
	if err != nil {
		t.Fatalf("Test_AnkaDB LoadConfig err %v", err)

		return
	}

	anka, err := NewAnkaDB(cfg, &BaseDBLogic{})
	if err != nil {
		t.Fatalf("Test_AnkaDB NewAnkaDB err %v", err)

		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	anka.RegEventFunc(EventOnStarted, func(ctxf context.Context, ankaf AnkaDB) error {
		for i := 0; i < 100; i++ {
			err := ankaf.Set(ctxf, "test001-1", fmt.Sprintf("test001-1-%d", i), []byte(fmt.Sprintf("test001-1-value-%d", i)))
			if err != nil {
				t.Fatalf("Test_AnkaDB Set err %v", err)

				return nil
			}

			err = ankaf.Set(ctxf, "test001-2", fmt.Sprintf("test001-2-%d", i), []byte(fmt.Sprintf("test001-2-value-%d", i)))
			if err != nil {
				t.Fatalf("Test_AnkaDB Set err %v", err)

				return nil
			}

			err = ankaf.Set(ctxf, "test001-3", fmt.Sprintf("test001-2-%d", i), []byte(fmt.Sprintf("test001-2-value-%d", i)))
			if err != ErrNotFoundDB {
				t.Fatalf("Test_AnkaDB Set err")

				return nil
			}
		}

		for i := 0; i < 100; i++ {
			val, err := ankaf.Get(ctxf, "test001-3", fmt.Sprintf("test001-1-%d", i))
			if err != ErrNotFoundDB {
				t.Fatalf("Test_AnkaDB Get err")

				return nil
			}

			val, err = ankaf.Get(ctxf, "test001-1", fmt.Sprintf("test001-1-%d", i))
			if err != nil {
				t.Fatalf("Test_AnkaDB Get err %v", err)

				return nil
			}

			if string(val) != fmt.Sprintf("test001-1-value-%d", i) {
				t.Fatalf("Test_AnkaDB Get fail")

				return nil
			}

			val, err = ankaf.Get(ctxf, "test001-1", fmt.Sprintf("test001-2-%d", i))
			if err != ErrNotFoundKey {
				t.Fatalf("Test_AnkaDB Get err")

				return nil
			}

			val, err = ankaf.Get(ctxf, "test001-2", fmt.Sprintf("test001-2-%d", i))
			if err != nil {
				t.Fatalf("Test_AnkaDB Get err %v", err)

				return nil
			}

			if string(val) != fmt.Sprintf("test001-2-value-%d", i) {
				t.Fatalf("Test_AnkaDB Get fail")

				return nil
			}

			val, err = ankaf.Get(ctxf, "test001-2", fmt.Sprintf("test001-1-%d", i))
			if err != ErrNotFoundKey {
				t.Fatalf("Test_AnkaDB Get err")

				return nil
			}
		}

		cancel()

		return nil
	})

	err = anka.Start(ctx)
	if err != nil {
		t.Fatalf("Test_AnkaDB Start err %v", err)

		return
	}

	t.Logf("Test_AnkaDB OK")
}
