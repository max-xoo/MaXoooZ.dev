package router

import (
	"MaXoooZ.dev/router/google"
	"MaXoooZ.dev/router/sfw"
	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/mcsrvstatus/{srv}", CreateRoute).Methods("GET")
	router.HandleFunc("/api/google/news", google.CreateNewsRoute).Methods("GET")

	router.HandleFunc("/api/sfw/cat", sfw.CreateCatRoute).Methods("GET")
	router.HandleFunc("/api/sfw/dog", sfw.CreateDogRoute).Methods("GET")

	return router
}
