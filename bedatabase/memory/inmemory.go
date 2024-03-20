package memory

import (
	"errors"
	"sync"
)

// InMemoryDatabase is an in-memory key-value database.
type InMemoryDatabase struct {
	data map[int]interface{}
	mu   sync.RWMutex
}

// NewInMemoryDatabase creates a new instance of InMemoryDatabase.
func NewInMemoryDatabase() *InMemoryDatabase {
	return &InMemoryDatabase{
		data: make(map[int]interface{}),
	}
}

// Connect connects to the in-memory database (no operation for in-memory implementation).
func (db *InMemoryDatabase) Connect() error {
	return nil
}

// Disconnect disconnects from the in-memory database (no operation for in-memory implementation).
func (db *InMemoryDatabase) Disconnect() error {
	return nil
}

// Insert inserts data into the in-memory database.
func (db *InMemoryDatabase) Insert(data interface{}) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	id := len(db.data) + 1
	db.data[id] = data
	return nil
}

// Update updates data in the in-memory database.
func (db *InMemoryDatabase) Update(id int, data interface{}) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	if _, exists := db.data[id]; !exists {
		return errors.New("record not found")
	}
	db.data[id] = data
	return nil
}

// Delete deletes data from the in-memory database.
func (db *InMemoryDatabase) Delete(id int) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	if _, exists := db.data[id]; !exists {
		return errors.New("record not found")
	}
	delete(db.data, id)
	return nil
}

// Query performs a query on the in-memory database (no operation for in-memory implementation).
func (db *InMemoryDatabase) Query(query string, args ...interface{}) (*map[interface{}]interface{}, error) {
	result := make(map[interface{}]interface{})
	for k, v := range db.data {
		result[k] = v
	}
	return &result, nil
}
