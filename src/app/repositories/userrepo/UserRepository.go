package userrepo

import (
	"simple-api/src/app/core/baserepo"

	"gopkg.in/mgo.v2"
)

//Repository is
type Repository struct {
	baserepo.RepositoryBase
	Schema *mgo.Collection
}
