package router

import (
	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/mcsrvstatus/{srv}", CreateRoute).Methods("GET")
	return router
}
