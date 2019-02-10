package database

import (
	"bytes"
	"strconv"
	"sync"
	"testing"
)

func Test_RocksDB(t *testing.T) {
	rdb, err := NewAnkaRDB("../test/rocksdb_test")
	if err != nil {
		t.Fatalf("Test_RocksDB NewAnkaRDB %v", err)

		return
	}

	if rdb == nil {
		t.Fatalf("Test_RocksDB NewAnkaRDB rdb is nil")

		return
	}

	if rdb.GetPath() != "../test/rocksdb_test" {
		t.Fatalf("Test_RocksDB ankaRDB.GetPath err %v", rdb.GetPath())

		return
	}

	hasnokey, err := rdb.Has([]byte("testkey"))
	if err != nil {
		t.Fatalf("Test_RocksDB ankaRDB.Has err %v", err)

		return
	}

	if hasnokey {
		t.Fatalf("Test_RocksDB ankaRDB.Has err")

		return
	}

	valnokey, err := rdb.Get([]byte("testkey"))
	if err != ErrNotFound {
		t.Fatalf("Test_RocksDB ankaRDB.Get err %v", err)

		return
	}

	if valnokey != nil {
		t.Fatalf("Test_RocksDB ankaRDB.Get err %v", valnokey)

		return
	}

	it := rdb.NewIterator()
	if it.Error() != nil {
		t.Fatalf("Test_RocksDB ankaRDB.NewIterator err %v", it.Error())

		return
	}

	if it.Next() {
		t.Fatalf("Test_RocksDB ankaRDB.NewIterator empty err")

		return
	}

	err = rdb.Put([]byte("testkey"), []byte("testkey value"))
	if err != nil {
		t.Fatalf("Test_RocksDB ankaRDB.Put err %v", err)

		return
	}

	hasnokey, err = rdb.Has([]byte("testkey"))
	if err != nil {
		t.Fatalf("Test_RocksDB ankaRDB.Has after Put err %v", err)

		return
	}

	if !hasnokey {
		t.Fatalf("Test_RocksDB ankaRDB.Has after Put err")

		return
	}

	valnokey, err = rdb.Get([]byte("testkey"))
	if err != nil {
		t.Fatalf("Test_RocksDB ankaRDB.Get after Put err %v", err)

		return
	}

	if !bytes.Equal(valnokey, []byte("testkey value")) {
		t.Fatalf("Test_RocksDB ankaRDB.Get value err %v", string(valnokey))

		return
	}

	err = rdb.Delete([]byte("testkey"))
	if err != nil {
		t.Fatalf("Test_RocksDB ankaRDB.Delete err %v", err)

		return
	}

	hasnokey, err = rdb.Has([]byte("testkey"))
	if err != nil {
		t.Fatalf("Test_RocksDB ankaRDB.Has after Delete err %v", err)

		return
	}

	if hasnokey {
		t.Fatalf("Test_RocksDB ankaRDB.Has after Delete err")

		return
	}

	err = rdb.Put([]byte("testkey"), []byte("reset testkey value"))
	if err != nil {
		t.Fatalf("Test_RocksDB ankaRDB.Put reset err %v", err)

		return
	}

	hasnokey, err = rdb.Has([]byte("testkey"))
	if err != nil {
		t.Fatalf("Test_RocksDB ankaRDB.Has after reset err %v", err)

		return
	}

	if !hasnokey {
		t.Fatalf("Test_RocksDB ankaRDB.Has after reset err")

		return
	}

	valnokey, err = rdb.Get([]byte("testkey"))
	if err != nil {
		t.Fatalf("Test_RocksDB ankaRDB.Get after reset err %v", err)

		return
	}

	if !bytes.Equal(valnokey, []byte("reset testkey value")) {
		t.Fatalf("Test_RocksDB ankaRDB.Get reset value err %v", string(valnokey))

		return
	}

	err = rdb.Put([]byte("testkey"), []byte("rewrite testkey value"))
	if err != nil {
		t.Fatalf("Test_RocksDB ankaRDB.Put rewrite err %v", err)

		return
	}

	hasnokey, err = rdb.Has([]byte("testkey"))
	if err != nil {
		t.Fatalf("Test_RocksDB ankaRDB.Has after rewrite err %v", err)

		return
	}

	if !hasnokey {
		t.Fatalf("Test_RocksDB ankaRDB.Has after rewrite err")

		return
	}

	valnokey, err = rdb.Get([]byte("testkey"))
	if err != nil {
		t.Fatalf("Test_RocksDB ankaRDB.Get after rewrite err %v", err)

		return
	}

	if !bytes.Equal(valnokey, []byte("rewrite testkey value")) {
		t.Fatalf("Test_RocksDB ankaRDB.Get rewrite value err %v", string(valnokey))

		return
	}

	for i := 0; i < 10; i++ {
		err = rdb.Put([]byte("testkey:"+strconv.Itoa(i)), []byte("testkey:"+strconv.Itoa(i)+" value"))
		if err != nil {
			t.Fatalf("Test_RocksDB ankaRDB.Put key:%v err %v", i, err)

			return
		}
	}

	it = rdb.NewIterator()
	if it.Error() != nil {
		t.Fatalf("Test_RocksDB ankaRDB.NewIterator err %v", it.Error())

		return
	}

	nums := 0
	for {
		if it.Valid() {
			nums++

			// t.Logf("Test_RocksDB %v - %v", string(it.Key()), string(it.Value()))
		}

		if !it.Next() {
			break
		}
	}

	if nums != 11 {
		t.Fatalf("Test_RocksDB ankaRDB.NewIterator nums err %v", nums)

		return
	}

	it.Release()

	it = rdb.NewIteratorWithPrefix([]byte("testkey:"))
	if it.Error() != nil {
		t.Fatalf("Test_RocksDB ankaRDB.NewIteratorWithPrefix err %v", it.Error())

		return
	}

	nums = 0
	for {
		if it.Valid() {
			nums++

			//t.Logf("Test_RocksDB %v - %v", string(it.Key()), string(it.Value()))
		}

		if !it.Next() {
			break
		}
	}

	if nums != 10 {
		t.Fatalf("Test_RocksDB ankaRDB.NewIteratorWithPrefix nums err %v", nums)

		return
	}

	it.Release()

	t.Logf("Test_RocksDB OK")
}

func Test_RocksDBAsync(t *testing.T) {
	rdb, err := NewAnkaRDB("../test/rocksdb_testasync")
	if err != nil {
		t.Fatalf("Test_RocksDBAsync NewAnkaRDB %v", err)

		return
	}

	if rdb == nil {
		t.Fatalf("Test_RocksDBAsync NewAnkaRDB rdb is nil")

		return
	}

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(index int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				// t.Logf(strconv.Itoa(index) + ":" + strconv.Itoa(j))

				err := rdb.Put([]byte("testkeyasync:"+strconv.Itoa(index)+":"+strconv.Itoa(j)),
					[]byte(strconv.Itoa(index)+":"+strconv.Itoa(j)))
				if err != nil {
					t.Fatalf("Test_RocksDBAsync ankaRDB.Put err %v", err)

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

				_, err := rdb.Get([]byte("testkeyasync:" + strconv.Itoa(index) + ":" + strconv.Itoa(j)))
				if err != nil {
					t.Fatalf("Test_RocksDBAsync ankaRDB.Get err %v", err)

					return
				}

				err = rdb.Put([]byte("testkeyasync:"+strconv.Itoa(9-index)+":"+strconv.Itoa(j)),
					[]byte(strconv.Itoa(index)+":"+strconv.Itoa(j)))
				if err != nil {
					t.Fatalf("Test_RocksDBAsync ankaRDB.Put err %v", err)

					return
				}
			}
		}(i)
	}

	wg.Wait()

	t.Logf("Test_RocksDBAsync OK")
}
