package google

import (
	"MaXoooZ.dev/utils"
	"encoding/json"
	"encoding/xml"
	"net/http"
	"strings"
)

func CreateNewsRoute(res http.ResponseWriter, req *http.Request) {
	search := req.FormValue("search")
	lang := req.FormValue("lang")

	i := strings.Index(search, " ")

	if search == "" || i > -1 {
		utils.NewError(400, "Your request is not valid")
		return
	}
	data, err := utils.GetNewsXML(search, lang)

	if err != nil {
		utils.NewError(500, "Internal Server Error")
		return
	}
	var rss utils.Rss
	xml.Unmarshal(data, &rss)

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(200)
	json.NewEncoder(res).Encode(rss.Channel)
}
