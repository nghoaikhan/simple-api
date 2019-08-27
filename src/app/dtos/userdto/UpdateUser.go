package userdto

import (
	"gopkg.in/mgo.v2/bson"
)

type UpdateDto struct {
	Username string   `json:"username,omitempty"`
	Password string   `json:"password,omitempty"`
	FullName string   `json:"fullname,omitemptty"`
	Role     string   `json:"role,omitempty"`
	Account  []string `json:"account,omitempty"`
}

func (b *UpdateDto) ToBsonMap() (bson.M, error) {
	bt, err := bson.Marshal(b)
	if err != nil {
		return nil, err
	}
	var bs bson.M
	err = bson.Unmarshal(bt, &bs)
	if err != nil {
		return nil, err
	}
	return bs, nil
}
