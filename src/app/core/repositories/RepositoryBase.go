package repositories

import (
	"gopkg.in/mgo.v2"
)

type RepositoryBase struct {
	model *mgo.Collection
}

func (re *RepositoryBase) Init(model *mgo.Collection) {
	re.model = model
}

func (re *RepositoryBase) Create() {

}
