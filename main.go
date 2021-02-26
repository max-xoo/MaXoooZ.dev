package main

import (
	"MaXoooZ.dev/router"
	"log"
	"net/http"
)

func main() {
	handler := router.Init()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	log.Println("Web server started on port 8080")
	log.Fatal(srv.ListenAndServe())
}
