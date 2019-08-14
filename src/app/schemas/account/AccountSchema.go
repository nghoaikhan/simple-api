package account

import (
	"gopkg.in/mgo.v2"

	"simple-api/src/app/core/common/fieldnames"
	"simple-api/src/app/core/utils"
)

var (
	Schema  *mgo.Collection
	session *mgo.Session
)

func init() {
	session = utils.GetSession().Copy()
	Schema = session.DB(utils.AppConfig.DBName).C(fieldnames.Account.Collection)
}

func CloseSession() {
	session.Close()
}
