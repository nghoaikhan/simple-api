package main

import (
	"fmt"
	"simple-api/src/app/core/common/fieldnames"
)

func main() {
	a := fieldnames.GetUser()
	fmt.Println(a.ID)
}
