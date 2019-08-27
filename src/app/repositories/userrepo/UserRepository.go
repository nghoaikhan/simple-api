package userrepo

import (
	"reflect"
	"simple-api/src/app/core/baserepo"

	"gopkg.in/mgo.v2"
)

//Repository is
type Repository struct {
	baserepo.RepositoryBase
}

func NewRepo(schema *mgo.Collection, session *mgo.Session, modelType reflect.Type) *Repository {
	var r Repository
	r.Schema = schema
	r.Session = session
	r.ModelType = modelType
	return &r
}
