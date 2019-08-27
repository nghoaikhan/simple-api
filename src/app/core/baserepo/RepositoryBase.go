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
func (re *RepositoryBase) FindByID(_id string) (result interface{}, err error) {
	tmp := reflect.New(re.ModelType).Interface()
	err = re.Schema.FindId(bson.ObjectIdHex(_id)).One(&tmp)
	result = tmp
	return
}

// Find is a
func (re *RepositoryBase) Find() (result interface{}, err error) {
	err = re.Schema.Find(nil).All(&result)
	return
}

// UpdateByID is a
func (re *RepositoryBase) UpdateByID(_id string, update updateDto) (err error) {
	err = re.Schema.UpdateId(bson.ObjectIdHex(_id), update)
	return
}

// Delete is a
func (re *RepositoryBase) Delete(_id string) (err error) {
	err = re.Schema.Remove(bson.M{"_id": bson.ObjectIdHex(_id)})
	return

}
