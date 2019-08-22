package user

import (
	"gopkg.in/mgo.v2/bson"
)

// Model is a
type Model struct {
	ID       bson.ObjectId `bson:"_id, omitempty"`
	Username string        `json:"username"`
	Password string        `json:"password"`
}
