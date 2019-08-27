package main

import (
	"net/http"
	"simple-api/src/app/core/utils"
	"simple-api/src/app/core/utils/mongo"
	"simple-api/src/app/routers"
)

func main() {
	utils.InitConfig()
	mongo.CreateDbSession()
	router := routers.InitRoutes()
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	server.ListenAndServe()

}
