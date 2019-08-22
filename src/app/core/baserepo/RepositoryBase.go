package baserepo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// RepositoryBase is a
type RepositoryBase struct {
	schema *mgo.Collection
}

// Init is a
func (re *RepositoryBase) Init(schema *mgo.Collection) {
	re.schema = schema

}

// Create is a
func (re *RepositoryBase) Create(user interface{}) error {
	err := re.schema.Insert(user)
	return err
}

// FindByID is a
func (re *RepositoryBase) FindByID(_id bson.ObjectId) (result interface{}, err error) {
	err = re.schema.Find(bson.M{"_id": _id}).One(&result)
	return
}

// Find is a
func (re *RepositoryBase) Find() (result []interface{}) {
	iter := re.schema.Find(nil).Iter()
	var doc interface{}
	for iter.Next(&doc) {
		result = append(result, doc)
	}
	return
}

// UpdateByID is a
func (re *RepositoryBase) UpdateByID(selector interface{}, update interface{}) (result interface{}) {

	return result
}

// Delete is a
func (re *RepositoryBase) Delete(_id bson.ObjectId) (err error) {
	err = re.schema.Remove(bson.M{"_id": _id})
	return

}
