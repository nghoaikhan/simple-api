package repositories

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// RepositoryBase is a
type RepositoryBase struct {
	model *mgo.Collection
}

// Init is a
func (re *RepositoryBase) Init(model *mgo.Collection) {
	re.model = model

}

// Create is a
func (re *RepositoryBase) Create(user interface{}) error {
	err := re.model.Insert(user)
	return err
}

// GetByID is a
func (re *RepositoryBase) FindById(_id bson.ObjectId) (result interface{}, err error) {
	err = re.model.Find(bson.M{"_id": _id}).One(&result)
	return
}

// GetAll is a
func (re *RepositoryBase) Find() (result []interface{}) {
	iter := re.model.Find(nil).Iter()
	var doc interface{}
	for iter.Next(&doc) {
		result = append(result, doc)
	}
	return
}

// Update is a
func (re *RepositoryBase) UpdateById(selector interface{}, update interface{}) (result interface{}) {

	return result
}

// Delete is a
func (re *RepositoryBase) Delete(_id bson.ObjectId) (err error) {
	err = re.model.Remove(bson.M{"_id": _id})
	return

}
