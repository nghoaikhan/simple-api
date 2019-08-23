package main

import (
	"fmt"
	"simple-api/src/app/core/utils"
	"simple-api/src/app/core/utils/mongo"
	"simple-api/src/ioc"
)

func main() {
	utils.InitConfig()
	mongo.CreateDbSession()
	userRepo := ioc.InitUserRepo()

	fmt.Println(userRepo.Find())

}
