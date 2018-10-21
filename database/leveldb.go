package database

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/iterator"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"
)

// ankaLDB - anka levelDB
type ankaLDB struct {
	dbpath string      // filename for reporting
	db     *leveldb.DB // LevelDB instance
}

// NewAnkaLDB - returns a ankaDB wrapped object.
func NewAnkaLDB(dbpath string, cache int, handles int) (Database, error) {
	// Ensure we have some minimal caching and file guarantees
	if cache < 16 {
		cache = 16
	}
	if handles < 16 {
		handles = 16
	}

	// Open the db and recover any potential corruptions
	db, err := leveldb.OpenFile(dbpath, &opt.Options{
		OpenFilesCacheCapacity: handles,
		BlockCacheCapacity:     cache / 2 * opt.MiB,
		WriteBuffer:            cache / 4 * opt.MiB, // Two of these are used internally
		Filter:                 filter.NewBloomFilter(10),
	})
	if _, corrupted := err.(*errors.ErrCorrupted); corrupted {
		db, err = leveldb.RecoverFile(dbpath, nil)
	}
	// (Re)check for errors and abort if opening of the db failed
	if err != nil {
		return nil, err
	}

	return &ankaLDB{
		dbpath: dbpath,
		db:     db,
	}, nil
}

// GetPath - returns the path to the database directory
func (db *ankaLDB) GetPath() string {
	return db.dbpath
}

// Put - put the key / value
func (db *ankaLDB) Put(key []byte, value []byte) error {
	return db.db.Put(key, value, nil)
}

// Has - is exist key
func (db *ankaLDB) Has(key []byte) (bool, error) {
	return db.db.Has(key, nil)
}

// Get - returns the given key if it's present
func (db *ankaLDB) Get(key []byte) ([]byte, error) {
	dat, err := db.db.Get(key, nil)
	if err != nil {
		return nil, err
	}

	return dat, nil
}

// Delete - deletes the key from the queue and database
func (db *ankaLDB) Delete(key []byte) error {
	return db.db.Delete(key, nil)
}

// NewIterator - new iterator
func (db *ankaLDB) NewIterator() iterator.Iterator {
	return db.db.NewIterator(nil, nil)
}

// NewIteratorWithPrefix - new iterator with prefix
func (db *ankaLDB) NewIteratorWithPrefix(prefix []byte) iterator.Iterator {
	return db.db.NewIterator(util.BytesPrefix(prefix), nil)
}

// Close - close database
func (db *ankaLDB) Close() {
	err := db.db.Close()
	if err == nil {
	} else {
	}
}

func (db *ankaLDB) LevelDB() *leveldb.DB {
	return db.db
}

// NewBatch - new batch
func (db *ankaLDB) NewBatch() Batch {
	return &ankaLDBBatch{
		db: db.db,
		b:  new(leveldb.Batch),
	}
}

// ankaLDBBatch - anka leveldb batch
type ankaLDBBatch struct {
	db   *leveldb.DB
	b    *leveldb.Batch
	size int
}

// Put - put key / value
func (b *ankaLDBBatch) Put(key, value []byte) error {
	b.b.Put(key, value)
	b.size += len(value)
	return nil
}

// Delete - delete key
func (b *ankaLDBBatch) Delete(key []byte) error {
	b.b.Delete(key)
	b.size++
	return nil
}

// Write - start batch
func (b *ankaLDBBatch) Write() error {
	return b.db.Write(b.b, nil)
}

// ValueSize - return contorl nums
func (b *ankaLDBBatch) ValueSize() int {
	return b.size
}

// Reset - reset batch
func (b *ankaLDBBatch) Reset() {
	b.b.Reset()
	b.size = 0
}
