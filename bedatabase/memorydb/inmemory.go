package memorydb

import (
	"errors"
	"sync"
)

// InMemoryDB is an in-memory key-value database.
type InMemoryDB struct {
	data map[int]interface{}
	mu   sync.RWMutex
}

// NewInMemoryDB creates a new instance of InMemoryDB.
func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		data: make(map[int]interface{}),
	}
}

// Connect connects to the in-memory database (no operation for in-memory implementation).
func (db *InMemoryDB) Connect() error {
	return nil
}


// Connect connects to the in-memory database (no operation for in-memory implementation).
func (db *InMemoryDB) Ping() error {
	return nil
}

// Disconnect disconnects from the in-memory database (no operation for in-memory implementation).
func (db *InMemoryDB) Disconnect() error {
	return nil
}

// Insert inserts data into the in-memory database.
func (db *InMemoryDB) Insert(data interface{}) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	id := len(db.data) + 1
	db.data[id] = data
	return nil
}

// Update updates data in the in-memory database.
func (db *InMemoryDB) Update(id int, data interface{}) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	if _, exists := db.data[id]; !exists {
		return errors.New("record not found")
	}
	db.data[id] = data
	return nil
}

// Delete deletes data from the in-memory database.
func (db *InMemoryDB) Delete(id int) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	if _, exists := db.data[id]; !exists {
		return errors.New("record not found")
	}
	delete(db.data, id)
	return nil
}

// Query performs a query on the in-memory database (no operation for in-memory implementation).
func (db *InMemoryDB) Query(query string, args ...interface{}) (*map[interface{}]interface{}, error) {
	result := make(map[interface{}]interface{})
	for k, v := range db.data {
		result[k] = v
	}
	return &result, nil
}
