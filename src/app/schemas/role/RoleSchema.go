package role

import (
	"log"

	"gopkg.in/mgo.v2"

	"simple-api/src/app/core/common/fieldnames"
	"simple-api/src/app/core/utils"
	mongoUtils "simple-api/src/app/core/utils/mongo"
)

var (
	Schema  *mgo.Collection
	session *mgo.Session
)

func init() {
	session = mongoUtils.GetSession().Copy()
	Schema = session.DB(utils.AppConfig.DBName).C(fieldnames.Role.Collection)
}

func CloseSession() {
	session.Close()
}

// AddIndexes make sure the ID of each document is unique
func AddIndexes() {
	var err error
	userIndex := mgo.Index{
		Key:        []string{"_id"},
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
