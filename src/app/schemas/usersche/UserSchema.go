package usersche

import (
	"log"
	"simple-api/src/app/core/utils"
	mongoUtils "simple-api/src/app/core/utils/mongo"

	"gopkg.in/mgo.v2"

	"simple-api/src/app/core/common/fieldnames"
)

var (
	// Schema
	Schema *mgo.Collection

	Session *mgo.Session
)

func init() {
	Session = mongoUtils.GetSession().Copy()
	Schema = Session.DB(utils.AppConfig.DBName).C(fieldnames.User.Collection)
}

func CloseSession() {
	Session.Close()
}

// AddIndexes make sure the ID of each document is unique
func AddIndexes() {
	var err error
	userIndex := mgo.Index{
		Key:        []string{"username"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}

	err = Schema.EnsureIndex(userIndex)
	defer CloseSession()
	if err != nil {
		log.Fatalf("[addIndexes]: %s\n", err)
	}
}
