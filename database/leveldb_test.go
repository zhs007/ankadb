package database

import (
	"bytes"
	"strconv"
	"sync"
	"testing"
)

func Test_LevelDB(t *testing.T) {
	ldb, err := NewAnkaLDB("../test/database_test", 16, 16)
	if err != nil {
		t.Fatalf("Test_LevelDB NewAnkaLDB %v", err)

		return
	}

	if ldb == nil {
		t.Fatalf("Test_LevelDB NewAnkaLDB ldb is nil")

		return
	}

	if ldb.GetPath() != "../test/database_test" {
		t.Fatalf("Test_LevelDB ankaLDB.GetPath err %v", ldb.GetPath())

		return
	}

	hasnokey, err := ldb.Has([]byte("testkey"))
	if err != nil {
		t.Fatalf("Test_LevelDB ankaLDB.Has err %v", err)

		return
	}

	if hasnokey {
		t.Fatalf("Test_LevelDB ankaLDB.Has err")

		return
	}

	valnokey, err := ldb.Get([]byte("testkey"))
	if err != ErrNotFound {
		t.Fatalf("Test_LevelDB ankaLDB.Get err %v", err)

		return
	}

	if valnokey != nil {
		t.Fatalf("Test_LevelDB ankaLDB.Get err %v", valnokey)

		return
	}

	it := ldb.NewIterator()
	if it.Error() != nil {
		t.Fatalf("Test_LevelDB ankaLDB.NewIterator err %v", it.Error())

		return
	}

	if it.Next() {
		t.Fatalf("Test_LevelDB ankaLDB.NewIterator empty err")

		return
	}

	err = ldb.Put([]byte("testkey"), []byte("testkey value"))
	if err != nil {
		t.Fatalf("Test_LevelDB ankaLDB.Put err %v", err)

		return
	}

	hasnokey, err = ldb.Has([]byte("testkey"))
	if err != nil {
		t.Fatalf("Test_LevelDB ankaLDB.Has after Put err %v", err)

		return
	}

	if !hasnokey {
		t.Fatalf("Test_LevelDB ankaLDB.Has after Put err")

		return
	}

	valnokey, err = ldb.Get([]byte("testkey"))
	if err != nil {
		t.Fatalf("Test_LevelDB ankaLDB.Get after Put err %v", err)

		return
	}

	if !bytes.Equal(valnokey, []byte("testkey value")) {
		t.Fatalf("Test_LevelDB ankaLDB.Get value err %v", string(valnokey))

		return
	}

	err = ldb.Delete([]byte("testkey"))
	if err != nil {
		t.Fatalf("Test_LevelDB ankaLDB.Delete err %v", err)

		return
	}

	hasnokey, err = ldb.Has([]byte("testkey"))
	if err != nil {
		t.Fatalf("Test_LevelDB ankaLDB.Has after Delete err %v", err)

		return
	}

	if hasnokey {
		t.Fatalf("Test_LevelDB ankaLDB.Has after Delete err")

		return
	}

	err = ldb.Put([]byte("testkey"), []byte("reset testkey value"))
	if err != nil {
		t.Fatalf("Test_LevelDB ankaLDB.Put reset err %v", err)

		return
	}

	hasnokey, err = ldb.Has([]byte("testkey"))
	if err != nil {
		t.Fatalf("Test_LevelDB ankaLDB.Has after reset err %v", err)

		return
	}

	if !hasnokey {
		t.Fatalf("Test_LevelDB ankaLDB.Has after reset err")

		return
	}

	valnokey, err = ldb.Get([]byte("testkey"))
	if err != nil {
		t.Fatalf("Test_LevelDB ankaLDB.Get after reset err %v", err)

		return
	}

	if !bytes.Equal(valnokey, []byte("reset testkey value")) {
		t.Fatalf("Test_LevelDB ankaLDB.Get reset value err %v", string(valnokey))

		return
	}

	err = ldb.Put([]byte("testkey"), []byte("rewrite testkey value"))
	if err != nil {
		t.Fatalf("Test_LevelDB ankaLDB.Put rewrite err %v", err)

		return
	}

	hasnokey, err = ldb.Has([]byte("testkey"))
	if err != nil {
		t.Fatalf("Test_LevelDB ankaLDB.Has after rewrite err %v", err)

		return
	}

	if !hasnokey {
		t.Fatalf("Test_LevelDB ankaLDB.Has after rewrite err")

		return
	}

	valnokey, err = ldb.Get([]byte("testkey"))
	if err != nil {
		t.Fatalf("Test_LevelDB ankaLDB.Get after rewrite err %v", err)

		return
	}

	if !bytes.Equal(valnokey, []byte("rewrite testkey value")) {
		t.Fatalf("Test_LevelDB ankaLDB.Get rewrite value err %v", string(valnokey))

		return
	}

	for i := 0; i < 10; i++ {
		err = ldb.Put([]byte("testkey:"+strconv.Itoa(i)), []byte("testkey:"+strconv.Itoa(i)+" value"))
		if err != nil {
			t.Fatalf("Test_LevelDB ankaLDB.Put key:%v err %v", i, err)

			return
		}
	}

	it = ldb.NewIterator()
	if it.Error() != nil {
		t.Fatalf("Test_LevelDB ankaLDB.NewIterator err %v", it.Error())

		return
	}

	nums := 0
	for {
		if it.Valid() {
			nums++

			t.Logf("Test_LevelDB %v - %v", string(it.Key()), string(it.Value()))
		}

		if !it.Next() {
			break
		}
	}

	if nums != 11 {
		t.Fatalf("Test_LevelDB ankaLDB.NewIterator nums err %v", nums)

		return
	}

	it.Release()

	it = ldb.NewIteratorWithPrefix([]byte("testkey:"))
	if it.Error() != nil {
		t.Fatalf("Test_LevelDB ankaLDB.NewIteratorWithPrefix err %v", it.Error())

		return
	}

	nums = 0
	for {
		if it.Valid() {
			nums++

			t.Logf("Test_LevelDB %v - %v", string(it.Key()), string(it.Value()))
		}

		if !it.Next() {
			break
		}
	}

	if nums != 10 {
		t.Fatalf("Test_LevelDB ankaLDB.NewIteratorWithPrefix nums err %v", nums)

		return
	}

	it.Release()

	t.Logf("Test_LevelDB OK")
}

func Test_LevelDBAsync(t *testing.T) {
	ldb, err := NewAnkaLDB("../test/database_testasync", 16, 16)
	if err != nil {
		t.Fatalf("Test_LevelDBAsync NewAnkaLDB %v", err)

		return
	}

	if ldb == nil {
		t.Fatalf("Test_LevelDBAsync NewAnkaLDB ldb is nil")

		return
	}

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(index int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				// t.Logf(strconv.Itoa(index) + ":" + strconv.Itoa(j))

				err := ldb.Put([]byte("testkeyasync:"+strconv.Itoa(index)+":"+strconv.Itoa(j)),
					[]byte(strconv.Itoa(index)+":"+strconv.Itoa(j)))
				if err != nil {
					t.Fatalf("Test_LevelDBAsync ankaLDB.Put err %v", err)

					return
				}
			}
		}(i)
	}

	wg.Wait()

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(index int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				// t.Logf(strconv.Itoa(index) + ":" + strconv.Itoa(j))

				_, err := ldb.Get([]byte("testkeyasync:" + strconv.Itoa(index) + ":" + strconv.Itoa(j)))
				if err != nil {
					t.Fatalf("Test_LevelDBAsync ankaLDB.Get err %v", err)

					return
				}

				err = ldb.Put([]byte("testkeyasync:"+strconv.Itoa(9-index)+":"+strconv.Itoa(j)),
					[]byte(strconv.Itoa(index)+":"+strconv.Itoa(j)))
				if err != nil {
					t.Fatalf("Test_LevelDBAsync ankaLDB.Put err %v", err)

					return
				}
			}
		}(i)
	}

	wg.Wait()

	t.Logf("Test_LevelDBAsync OK")
}
