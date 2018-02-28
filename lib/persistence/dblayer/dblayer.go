package dblayer

import (
	"github.com/user/rest-api/lib/persistence"
	"github.com/user/rest-api/lib/persistence/mongolayer"
)

type DBTYPE string

const (
	MONGODB DBTYPE = "mongodb"
	DYNAMODB DBTYPE = "dynamodb"
)

func NewPersistenceLayer(options DBTYPE, connection string) (persistence.DatabaseHandler, error) {
	switch options {
	case MONGODB:
		return mongolayer.NewMongoDBLayer(connection)
	}
	return nil, nil
}