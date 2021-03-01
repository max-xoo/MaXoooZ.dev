package main

import (
	"MaXoooZ.dev/router"
	"MaXoooZ.dev/utils"
	"log"
	"net/http"
)

func main() {
	utils.GetNewsXML("dev", "fr")

	handler := router.Init()

	srv := &http.Server{
		Addr:    ":80",
		Handler: handler,
	}

	log.Println("Web server started on port 80")
	log.Fatal(srv.ListenAndServe())
}
