package main

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

type user struct {
	NamePro string
	Email   string
}

func main() {
	newProj := &user{
		NamePro: "foo",
		Email:   "khan",
	}
	ms, _ := bson.Marshal(newProj)
	var bm bson.M
	_ = bson.Unmarshal(ms, &bm)
	bm2 := bson.M{"Name": "foo", "Email": "khan"}
	fmt.Println(bm)
	fmt.Println(bm2)
}
