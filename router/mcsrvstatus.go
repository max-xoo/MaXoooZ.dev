package router

import (
	"MaXoooZ.dev/query"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func CreateRoute(res http.ResponseWriter, req *http.Request) {
	srv := mux.Vars(req)["srv"]

	request := query.NewRequest()
	request.Connect(srv)
	data, _ := request.Full()

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(200)
	json.NewEncoder(res).Encode(data)
}
