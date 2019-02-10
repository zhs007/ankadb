package database

import (
	"github.com/tecbot/gorocksdb"
)

// ankaRDB - anka RocksDB
type ankaRDB struct {
	dbpath string        // filename for reporting
	db     *gorocksdb.DB // RocksDB instance
	ro     *gorocksdb.ReadOptions
	wo     *gorocksdb.WriteOptions
	itro   *gorocksdb.ReadOptions
}

// NewAnkaRDB - returns a ankaDB wrapped object.
func NewAnkaRDB(dbpath string) (Database, error) {
	bbto := gorocksdb.NewDefaultBlockBasedTableOptions()
	bbto.SetBlockCache(nil) //gorocksdb.NewLRUCache(3 << 30))
	opts := gorocksdb.NewDefaultOptions()
	opts.SetBlockBasedTableFactory(bbto)
	opts.SetCreateIfMissing(true)
	db, err := gorocksdb.OpenDb(opts, dbpath)
	if err != nil {
		return nil, err
	}

	itro := gorocksdb.NewDefaultReadOptions()
	itro.SetFillCache(false)

	return &ankaRDB{
		dbpath: dbpath,
		db:     db,
		ro:     gorocksdb.NewDefaultReadOptions(),
		wo:     gorocksdb.NewDefaultWriteOptions(),
		itro:   itro,
	}, nil
}

// GetPath - returns the path to the database directory
func (rdb *ankaRDB) GetPath() string {
	return rdb.dbpath
}

// Put - put the key / value
func (rdb *ankaRDB) Put(key []byte, value []byte) error {
	return rdb.db.Put(rdb.wo, key, value)
}

// Has - is exist key
func (rdb *ankaRDB) Has(key []byte) (bool, error) {
	value, err := rdb.db.GetBytes(rdb.ro, key)
	if err != nil {
		return false, err
	}

	if value == nil {
		return false, nil
	}

	return true, nil
}

// Get - returns the given key if it's present
func (rdb *ankaRDB) Get(key []byte) ([]byte, error) {
	value, err := rdb.db.GetBytes(rdb.ro, key)
	if value == nil {
		return nil, ErrNotFound
	}

	return value, err
}

// Delete - deletes the key from the queue and database
func (rdb *ankaRDB) Delete(key []byte) error {
	return rdb.db.Delete(rdb.wo, key)
}

// NewIterator - new iterator
func (rdb *ankaRDB) NewIterator() Iterator {
	it := rdb.db.NewIterator(rdb.itro)
	return &iteratorRDB{
		iter: it,
	}
}

// NewIteratorWithPrefix - new iterator with prefix
func (rdb *ankaRDB) NewIteratorWithPrefix(prefix []byte) Iterator {
	it := rdb.db.NewIterator(rdb.itro)
	return &iteratorRDB{
		iter: it,
	}
}

// Close - close database
func (rdb *ankaRDB) Close() {
	rdb.db.Close()
	rdb.ro.Destroy()
	rdb.wo.Destroy()
	rdb.itro.Destroy()
}

// NewBatch - new batch
func (rdb *ankaRDB) NewBatch() Batch {
	wb := gorocksdb.NewWriteBatch()

	return &batchRDB{
		rdb: rdb,
		wb:  wb,
	}
}

type batchRDB struct {
	rdb *ankaRDB
	wb  *gorocksdb.WriteBatch
}

// Put - put key / value
func (batch *batchRDB) Put(key, value []byte) error {
	batch.wb.Put(key, value)

	return nil
}

// Delete - delete key
func (batch *batchRDB) Delete(key []byte) error {
	batch.wb.Delete(key)

	return nil
}

// Write - start batch
func (batch *batchRDB) Write() error {
	return batch.rdb.db.Write(batch.rdb.wo, batch.wb)
}

// ValueSize - return contorl nums
func (batch *batchRDB) ValueSize() int {
	return batch.wb.Count()
}

// Reset - reset batch
func (batch *batchRDB) Reset() {
	batch.wb.Clear()
}

// Release - release
func (batch *batchRDB) Release() {
	batch.wb.Destroy()
}

type iteratorRDB struct {
	iter *gorocksdb.Iterator
}

// First - seek to first
func (it *iteratorRDB) First() bool {
	it.iter.SeekToFirst()
	return it.iter.Valid()
}

// Last - seek to last
func (it *iteratorRDB) Last() bool {
	it.iter.SeekToLast()
	return it.iter.Valid()
}

// Seek - seek
func (it *iteratorRDB) Seek(key []byte) bool {
	it.iter.Seek(key)
	return it.iter.Valid()
}

// Next - seek to next
func (it *iteratorRDB) Next() bool {
	it.iter.Next()
	return it.iter.Valid()
}

// Prev - seek to prev
func (it *iteratorRDB) Prev() bool {
	it.iter.Prev()
	return it.iter.Valid()
}

// Valid - is valid
func (it *iteratorRDB) Valid() bool {
	return it.iter.Valid()
}

// Error - get error
func (it *iteratorRDB) Error() error {
	return it.iter.Err()
}

// Key - get key
func (it *iteratorRDB) Key() []byte {
	k := it.iter.Key()
	defer k.Free()
	return k.Data()
}

// Value - get value
func (it *iteratorRDB) Value() []byte {
	v := it.iter.Value()
	defer v.Free()
	return v.Data()
}

// Release - release
func (it *iteratorRDB) Release() {
	it.iter.Close()
}
