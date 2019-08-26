package baserepo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type createDto interface{}
type updateDto interface{}

// RepositoryBase is a
type RepositoryBase struct {
	Schema  *mgo.Collection
	Session *mgo.Session
}

// Create is a
func (re *RepositoryBase) Create(user createDto) (err error) {
	err = re.Schema.Insert(user)
	return
}

// FindByID is a
func (re *RepositoryBase) FindByID(_id bson.ObjectId) (result interface{}, err error) {
	err = re.Schema.Find(bson.M{"_id": _id}).One(&result)
	return
}

// Find is a
func (re *RepositoryBase) Find() (result []interface{}) {
	iter := re.Schema.Find(nil).Iter()
	var doc interface{}
	for iter.Next(&doc) {
		result = append(result, doc)
	}
	return
}

// UpdateByID is a
func (re *RepositoryBase) UpdateByID(selector interface{}, update updateDto) (result interface{}) {

	return result
}

// Delete is a
func (re *RepositoryBase) Delete(_id bson.ObjectId) (err error) {
	err = re.Schema.Remove(bson.M{"_id": _id})
	return

}
