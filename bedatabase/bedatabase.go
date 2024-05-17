package bedatabase

import "bedatabase/memorydb"

// DatabaseInterface defines the methods that a database interface must implement.
type BeDatabase interface {
	Connect() error
	Ping() error
	Disconnect() error
	Insert(data interface{}) error
	Update(id int, data interface{}) error
	Delete(id int) error
	Query(query string, args ...interface{}) (*map[interface{}]interface{}, error)
}

var _ BeDatabase = &memorydb.InMemoryDB{}

// // NewInMemoryDatabase creates a new instance of InMemoryDatabase.
// func NewInMemoryDatabase() BeDatabase {
// 	return memorydb.NewInMemoryDB()
// }
