package main

import (
	"fmt"
	"net/http"
	"simple-api/src/app/core/utils"
	"simple-api/src/app/core/utils/mongo"
	"simple-api/src/app/routers"
	"simple-api/src/ioc"
)

func main() {
	utils.InitConfig()
	mongo.CreateDbSession()
	userserv := ioc.InitUserServ()
	fmt.Println(userserv.GetUsers())
	router := routers.InitRoutes()
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	server.ListenAndServe()

}
