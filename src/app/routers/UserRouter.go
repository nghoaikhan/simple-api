package routers

import (
	"simple-api/src/app/core/common/enums"
	"simple-api/src/ioc"

	"github.com/gorilla/mux"
)

func setUserRouter(router *mux.Router) *mux.Router {
	usercontrl := ioc.InitUserContrl()
	router.HandleFunc("/users", usercontrl.GetUsers).Methods(enums.GetHttpMethod().GET)
	router.HandleFunc("/users/{id}", usercontrl.GetUserByID).Methods(enums.GetHttpMethod().GET)
	router.HandleFunc("/users", usercontrl.CreateUser).Methods(enums.GetHttpMethod().POST)
	router.HandleFunc("/users/{id}", usercontrl.UpdateUserByID).Methods(enums.GetHttpMethod().PUT)
	router.HandleFunc("/users/{id}", usercontrl.DeleteByID).Methods(enums.GetHttpMethod().DELETE)
	return router
}
