package bedatabase

import "bedatabase/memory"

// DatabaseInterface defines the methods that a database interface must implement.
type BeDatabase interface {
	Connect() error
	Disconnect() error
	Insert(data interface{}) error
	Update(id int, data interface{}) error
	Delete(id int) error
	Query(query string, args ...interface{}) (*map[interface{}]interface{}, error)
}

var _ BeDatabase = &memory.InMemoryDatabase{}

// NewInMemoryDatabase creates a new instance of InMemoryDatabase.
func NewInMemoryDatabase() BeDatabase {
	return memory.NewInMemoryDatabase()
}
