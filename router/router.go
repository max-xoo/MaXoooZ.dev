package router

import (
	"MaXoooZ.dev/router/sfw"
	"MaXoooZ.dev/router/web"
	"MaXoooZ.dev/utils"
	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/mcsrvstatus/{srv}", CreateRoute).Methods("GET")
	router.HandleFunc("/api/google/news", web.CreateNewsRoute).Methods("GET")

	router.HandleFunc("/api/sfw/cat", sfw.CreateCatRoute).Methods("GET")
	router.HandleFunc("/api/sfw/dog", sfw.CreateDogRoute).Methods("GET")

	router.Use(utils.LoggingMiddleware)
	return router
}
