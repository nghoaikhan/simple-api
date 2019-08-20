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

func (re *RepositoryBase) Create(user interface{}) error {
	err := re.model.Insert(user)
	return err
}

func (re *RepositoryBase) Get() (result interface{}) {
	return result
}

func (re *RepositoryBase) GetAll() (result interface{}) {
	result = re.model
	return result
}

func (re *RepositoryBase) Update(updateUser interface{}) (result interface{}) {
	return result
}

func (re *RepositoryBase) Delete(_id interface{}) (result interface{}) {
	return result
}
