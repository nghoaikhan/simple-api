package routers

import (
	"simple-api/src/app/core/common/enums"
	"simple-api/src/ioc"

	"github.com/gorilla/mux"
)

func setUserRouter(router *mux.Router) *mux.Router {
	usercontrl := ioc.InitUserContrl()
	router.HandleFunc("/users", usercontrl.CreateUser).Methods(enums.GetHttpMethod().POST)
	return router
}
