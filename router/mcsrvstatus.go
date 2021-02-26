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
	explode := strings.Split(srv, ":")

	data, err := query.Query(explode[0], explode[1]) // GoQuery of seyz :) https://github.com/Seyz123/GoQuery

	if err != nil {
		utils.NewError(400, "The server are not online or not exist")
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(200)
	json.NewEncoder(res).Encode(data)
}
