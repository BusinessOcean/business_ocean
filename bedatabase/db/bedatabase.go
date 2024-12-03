package db

// DatabaseInterface defines the methods that a database interface must implement.
type IBeDatabase interface {
	Connect() error
	Ping() error
	Disconnect() error
	Insert(data interface{}) error
	Update(id int, data interface{}) error
	Delete(id int) error
	Query(query string, args ...interface{}) (*map[interface{}]interface{}, error)
}
