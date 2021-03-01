package router

import (
	"MaXoooZ.dev/query"
	"MaXoooZ.dev/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func CreateRoute(res http.ResponseWriter, req *http.Request) {
	srv := mux.Vars(req)["srv"]
	i := strings.Index(srv, ":")

	if srv == "" || i <= -1 {
		utils.NewError(400, "Your request is not valid")
		return
	}

	request := query.NewRequest()
	request.Connect(srv)
	data, _ := request.Full()

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(200)
	json.NewEncoder(res).Encode(data)
}
