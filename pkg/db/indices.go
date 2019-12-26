package db

import (
	"fmt"
	"github.com/globalsign/mgo"
)

// CreateDatabaseIndices creates required indices for all collection
func CreateDatabaseIndices() error {
	return createTodoIndices()
}

func createTodoIndices() error {
	collection := GetDatabaseCollection("todo")
	err := collection.EnsureIndex(mgo.Index{Key: []string{"id", "userId"}})
	if err != nil {
		return err
	}
	fmt.Println("Successfully created todo indices")
	return nil
}
