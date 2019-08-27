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
	newProj := user{
		NamePro: "foo",
		Email:   "khan",
	}
	arr := []user{newProj}
	var bm []bson.M
	var tmpBson bson.M
	for _, v := range arr {
		ms, _ := bson.Marshal(&v)
		_ = bson.Unmarshal(ms, &tmpBson)
		bm = append(bm, tmpBson)
	}

	fmt.Println(bm)

}
