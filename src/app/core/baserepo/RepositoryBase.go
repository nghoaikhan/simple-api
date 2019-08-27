package baserepo

import (
	"reflect"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type createDto interface{}
type updateDto interface{}

// RepositoryBase is a
type RepositoryBase struct {
	Schema    *mgo.Collection
	Session   *mgo.Session
	ModelType reflect.Type
}

// Create is a
func (re *RepositoryBase) Create(user createDto) (err error) {
	err = re.Schema.Insert(user)
	return
}

// FindByID is a
func (re *RepositoryBase) FindByID(id string) (result interface{}, err error) {

	err = re.Schema.FindId(bson.ObjectIdHex(id)).One(&result)
	return
}

// Find is a
func (re *RepositoryBase) Find() (result []interface{}, err error) {
	err = re.Schema.Find(nil).All(&result)
	return
}

// UpdateByID is a
func (re *RepositoryBase) UpdateByID(id string, update bson.M) (err error) {
	err = re.Schema.UpdateId(bson.ObjectIdHex(id), update)
	return
}

// Delete is a
func (re *RepositoryBase) DeleteByID(id string) (err error) {
	err = re.Schema.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return

}
