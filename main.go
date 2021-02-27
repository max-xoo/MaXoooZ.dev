package main

import (
	"MaXoooZ.dev/router"
	"log"
	"net/http"
)

func main() {
	handler := router.Init()

	srv := &http.Server{
		Addr:    ":80",
		Handler: handler,
	}

	log.Println("Web server started on port 80")
	log.Fatal(srv.ListenAndServe())
}
