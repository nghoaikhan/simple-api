package main

import (
	"fmt"
	"simple-api/src/app/core/common/fieldnames"
)

func main() {
	fmt.Println(fieldnames.ACCOUNT.AddOnFeatures)
	fmt.Println(fieldnames.USER.Password)
}
