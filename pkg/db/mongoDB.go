package db

import (
	"github.com/globalsign/mgo"
)

var connection *mgo.Database

// ConnectToDatabase returns mongo connection to database matching connection string and database name parameters
func ConnectToDatabase(connectionString string, database string) (*mgo.Database, error) {
	session, connectionError := mgo.Dial(connectionString)
	if connectionError != nil {
		return nil, connectionError
	}
	connection = session.DB(database)
	return connection, nil
}

// GetDatabaseCollection returns a mongo collection given a collection name param
func GetDatabaseCollection(collectionName string) *mgo.Collection {
	return connection.C(collectionName)
}

// GetDatabaseConnection returns existing database connection
func GetDatabaseConnection() *mgo.Database {
	return connection
}
