package main

import (
	"net/http"
	"simple-api/src/app/core/utils"
	"simple-api/src/app/core/utils/mongo"
	"simple-api/src/ioc"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Name      string
	Phone     string
	Timestamp time.Time
}

func main() {
	utils.InitConfig()
	mongo.CreateDbSession()
	usercontrl := ioc.InitUserContrl()
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/add", usercontrl.CreateUser).Methods("POST")

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	server.ListenAndServe()

}
