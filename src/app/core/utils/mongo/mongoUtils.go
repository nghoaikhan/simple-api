package mongo

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"

	"simple-api/src/app/core/utils"
)

var session *mgo.Session

// GetSession create the session to mongoDB
func GetSession() *mgo.Session {
	if session == nil {
		var err error
		session, err = mgo.DialWithInfo(&mgo.DialInfo{
			Addrs:    []string{utils.AppConfig.MongoDBHost},
			Username: utils.AppConfig.DBUser,
			Password: utils.AppConfig.DBPwd,
			Timeout:  60 * time.Second,
		})
		if err != nil {
			log.Fatalf("[GetSession]: %s\n", err)
		}
	}
	return session
}

// CreateDbSession create a DB Session
func CreateDbSession() {
	var err error
	session, err = mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{utils.AppConfig.MongoDBHost},
		Username: utils.AppConfig.DBUser,
		Password: utils.AppConfig.DBPwd,
		Timeout:  60 * time.Second,
	})
	if err != nil {
		log.Fatalf("[createDbSession]: %s\n", err)
	}
}
