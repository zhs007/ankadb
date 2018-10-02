package database

import "github.com/syndtr/goleveldb/leveldb/iterator"

// Database -
type Database interface {
	// Path returns the path to the database directory.
	GetPath() string

	// Put puts the given key / value to the queue
	Put(key []byte, value []byte) error

	Has(key []byte) (bool, error)

	// Get returns the given key if it's present.
	Get(key []byte) ([]byte, error)

	// Delete deletes the key from the queue and database
	Delete(key []byte) error

	NewIterator() iterator.Iterator

	NewIteratorWithPrefix(prefix []byte) iterator.Iterator

	// // NewIteratorWithPrefix returns a iterator to iterate over subset of database content with a particular prefix.
	// NewIteratorWithPrefix(prefix []byte) iterator.Iterator

	Close()

	NewBatch() Batch
}

// Batch -
type Batch interface {
	Put(key, value []byte) error

	Delete(key []byte) error

	Write() error

	ValueSize() int

	Reset()
}
