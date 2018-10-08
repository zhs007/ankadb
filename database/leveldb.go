package database

import (
	"go.uber.org/zap"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/iterator"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"

	"github.com/zhs007/jarviscore/log"
)

// const (
// 	writePauseWarningThrottler = 1 * time.Minute
// )

// var OpenFileLimit = 64

type ankaLDB struct {
	dbpath string      // filename for reporting
	db     *leveldb.DB // LevelDB instance

	// quitLock sync.Mutex      // Mutex protecting the quit channel access
	// quitChan chan chan error // Quit channel to stop the metrics collection before closing the database
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

	log.Info("NewJarvisDB", zap.Int("cache", cache), zap.Int("handles", handles))

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

// Path returns the path to the database directory.
func (db *ankaLDB) GetPath() string {
	return db.dbpath
}

// Put puts the given key / value to the queue
func (db *ankaLDB) Put(key []byte, value []byte) error {
	return db.db.Put(key, value, nil)
}

func (db *ankaLDB) Has(key []byte) (bool, error) {
	return db.db.Has(key, nil)
}

// Get returns the given key if it's present.
func (db *ankaLDB) Get(key []byte) ([]byte, error) {
	dat, err := db.db.Get(key, nil)
	if err != nil {
		return nil, err
	}
	return dat, nil
}

// Delete deletes the key from the queue and database
func (db *ankaLDB) Delete(key []byte) error {
	return db.db.Delete(key, nil)
}

func (db *ankaLDB) NewIterator() iterator.Iterator {
	return db.db.NewIterator(nil, nil)
}

// NewIteratorWithPrefix returns a iterator to iterate over subset of database content with a particular prefix.
func (db *ankaLDB) NewIteratorWithPrefix(prefix []byte) iterator.Iterator {
	return db.db.NewIterator(util.BytesPrefix(prefix), nil)
}

func (db *ankaLDB) Close() {
	// // Stop the metrics collection to avoid internal database races
	// db.quitLock.Lock()
	// defer db.quitLock.Unlock()

	// if db.quitChan != nil {
	// 	errc := make(chan error)
	// 	db.quitChan <- errc
	// 	if err := <-errc; err != nil {
	// 		log.Error("Metrics collection failed", "err", err)
	// 	}
	// 	db.quitChan = nil
	// }
	err := db.db.Close()
	if err == nil {
		log.Info("ankaLDB:Close")
	} else {
		log.Error("ankaLDB:Close", zap.String("err", err.Error()))
	}
}

func (db *ankaLDB) LDB() *leveldb.DB {
	return db.db
}

func (db *ankaLDB) NewBatch() Batch {
	return &ldbBatch{db: db.db, b: new(leveldb.Batch)}
}

type ldbBatch struct {
	db   *leveldb.DB
	b    *leveldb.Batch
	size int
}

func (b *ldbBatch) Put(key, value []byte) error {
	b.b.Put(key, value)
	b.size += len(value)
	return nil
}

func (b *ldbBatch) Delete(key []byte) error {
	b.b.Delete(key)
	b.size++
	return nil
}

func (b *ldbBatch) Write() error {
	return b.db.Write(b.b, nil)
}

func (b *ldbBatch) ValueSize() int {
	return b.size
}

func (b *ldbBatch) Reset() {
	b.b.Reset()
	b.size = 0
}

type table struct {
	db     Database
	prefix string
}

// NewTable returns a Database object that prefixes all keys with a given
// string.
func NewTable(db Database, prefix string) Database {
	return &table{
		db:     db,
		prefix: prefix,
	}
}

func (dt *table) GetPath() string {
	return dt.db.GetPath() + ":" + dt.prefix
}

func (dt *table) Put(key []byte, value []byte) error {
	return dt.db.Put(append([]byte(dt.prefix), key...), value)
}

func (dt *table) Has(key []byte) (bool, error) {
	return dt.db.Has(append([]byte(dt.prefix), key...))
}

func (dt *table) Get(key []byte) ([]byte, error) {
	return dt.db.Get(append([]byte(dt.prefix), key...))
}

func (dt *table) Delete(key []byte) error {
	return dt.db.Delete(append([]byte(dt.prefix), key...))
}

func (dt *table) NewIterator() iterator.Iterator {
	return dt.db.NewIteratorWithPrefix([]byte(dt.prefix))
}

// NewIteratorWithPrefix returns a iterator to iterate over subset of database content with a particular prefix.
func (dt *table) NewIteratorWithPrefix(prefix []byte) iterator.Iterator {
	return dt.db.NewIteratorWithPrefix(append([]byte(dt.prefix), prefix...))
}

func (dt *table) Close() {
	// Do nothing; don't close the underlying DB.
}

type tableBatch struct {
	batch  Batch
	prefix string
}

// NewTableBatch returns a Batch object which prefixes all keys with a given string.
func NewTableBatch(db Database, prefix string) Batch {
	return &tableBatch{db.NewBatch(), prefix}
}

func (dt *table) NewBatch() Batch {
	return &tableBatch{dt.db.NewBatch(), dt.prefix}
}

func (tb *tableBatch) Put(key, value []byte) error {
	return tb.batch.Put(append([]byte(tb.prefix), key...), value)
}

func (tb *tableBatch) Delete(key []byte) error {
	return tb.batch.Delete(append([]byte(tb.prefix), key...))
}

func (tb *tableBatch) Write() error {
	return tb.batch.Write()
}

func (tb *tableBatch) ValueSize() int {
	return tb.batch.ValueSize()
}

func (tb *tableBatch) Reset() {
	tb.batch.Reset()
}
