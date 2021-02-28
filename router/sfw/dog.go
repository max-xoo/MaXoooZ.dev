package sfw

import (
	"MaXoooZ.dev/utils"
	"encoding/json"
	"net/http"
)

func CreateDogRoute(res http.ResponseWriter, req *http.Request) {
	bytes, err := utils.GetIMG("sfw", "dog")

	if err != nil {
		bytes, _ := json.Marshal(err)

		res.WriteHeader(err.Code)
		_, _ = res.Write(bytes)
		return
	}

	res.Header().Set("Content-Type", "image/jpg")
	res.WriteHeader(200)
	_, _ = res.Write(bytes)
}
