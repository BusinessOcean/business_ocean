package dgraphdb

import (
	"beconst"
	"becore/belogger"
	"bedatabase/db"
	"context"
	"fmt"

	"github.com/dgraph-io/dgo/v230"
	"github.com/dgraph-io/dgo/v230/protos/api"
)

// NOTE:   GRAPHQL_ENDPOINT :  https://green-feather-41381579.ap-south-1.aws.cloud.dgraph.io/graphql
//         GRPC_ENDPOINT    :  green-feather-41381579.grpc.ap-south-1.aws.cloud.dgraph.io:443
//         API_TOKEN_CLIENT :  YjM3YzcyZDI3YjJmMTM3ODNjZjMzZmYzNmQwNTRlOWI=
//         API_TOKEN_ADMIN  :  M2YwY2E3OGVhNTQ3YzI0YTA3YTliYThmZDdjZDRlZjU=

var _ db.IBeDatabase = (*BeDgraphDB)(nil)

// BeDgraphDB is an in-memory key-value database.
type BeDgraphDB struct {
	*dgo.Dgraph
	logger *belogger.BeLogger
}

// NewBeDgraphDB creates a new instance of BeDgraphDB.
func NewBeDgraphDB(logger *belogger.BeLogger) *BeDgraphDB {
	return &BeDgraphDB{logger: logger}
}

// Connect connects to the in-memory database (no operation for in-memory implementation).
func (db *BeDgraphDB) Connect() error {
	db.logger.Info("Connecting to DgraphDB...")

	conn, err := dgo.DialCloud(beconst.DGRAPH_GRAPHQL_ENDPOINT, beconst.DGRAPH_ADMIN_API_TOKEN)
	// Check error
	if err != nil {
		return err
	}

	dgraphClient := dgo.NewDgraphClient(api.NewDgraphClient(conn))
	db.Dgraph = dgraphClient

	return nil
}

// Connect connects to the in-memory database (no operation for in-memory implementation).
func (db *BeDgraphDB) Ping() error {
	err := db.Ping()
	fmt.Println("Ping DgraphDB", err)
	return nil
}

// Disconnect disconnects from the in-memory database (no operation for in-memory implementation).
func (db *BeDgraphDB) Disconnect() error {

	err := db.Alter(context.Background(), &api.Operation{DropOp: api.Operation_ALL})
	return err
}

// Insert inserts data into the in-memory database.
func (db *BeDgraphDB) Insert(data interface{}) error {

	return nil
}

// Update updates data in the in-memory database.
func (db *BeDgraphDB) Update(id int, data interface{}) error {

	return nil
}

// Delete deletes data from the in-memory database.
func (db *BeDgraphDB) Delete(id int) error {

	return nil
}

// Query performs a query on the in-memory database (no operation for in-memory implementation).
func (db *BeDgraphDB) Query(query string, args ...interface{}) (*map[interface{}]interface{}, error) {

	return nil, nil
}
