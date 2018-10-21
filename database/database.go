package database

import "github.com/syndtr/goleveldb/leveldb/iterator"

// Database - key / value database
type Database interface {
	// GetPath - returns the path to the database directory
	GetPath() string
	// Put - put the key / value
	Put(key []byte, value []byte) error
	// Has - is exist key
	Has(key []byte) (bool, error)
	// Get - returns the given key if it's present
	Get(key []byte) ([]byte, error)
	// Delete - deletes the key from the queue and database
	Delete(key []byte) error
	// NewIterator - new iterator
	NewIterator() iterator.Iterator
	// NewIteratorWithPrefix - new iterator with prefix
	NewIteratorWithPrefix(prefix []byte) iterator.Iterator
	// Close - close database
	Close()
	// NewBatch - new batch
	NewBatch() Batch
}

// Batch - batch control
type Batch interface {
	// Put - put key / value
	Put(key, value []byte) error
	// Delete - delete key
	Delete(key []byte) error
	// Write - start batch
	Write() error
	// ValueSize - return contorl nums
	ValueSize() int
	// Reset - reset batch
	Reset()
}
