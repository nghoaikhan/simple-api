package usermod

import (
	"simple-api/src/app/dtos/userdto"

	"time"

	"gopkg.in/mgo.v2/bson"
)

// Model is a
type Model struct {
	ID        bson.ObjectId `bson:"_id"`
	Username  string        `json:"username"`
	Password  string        `json:"password"`
	Accounts  []string      `json:"accounts"`
	Role      string        `json:"role"`
	FullName  string        `json:"fullname,omitempty"`
	CreatedOn time.Time     `json:"createdOn"`
}

func (m *Model) InitFromCreateDto(dto userdto.CreateDto) {
	m.ID = bson.NewObjectId()
	m.Username = dto.Username
	m.Password = dto.Password
	m.CreatedOn = time.Now()

}
