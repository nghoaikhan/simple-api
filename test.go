package main

import (
	"fmt"

	"github.com/fatih/structs"
)

type user struct {
	name  string
	email string
}

func main() {
	var a interface{}
	a = &user{"Khan", "@"}
	fmt.Println(structs.Names(&user{}))

}
