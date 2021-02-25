package main

import (
	"github.com/henrylee2cn/faygo"
	"MaXoooZ.dev/query"
)

type Index struct {
	Mcsrvstatus string `param:"<in:path> <required> <desc:MCSRVSTATUS>"`
}

func (i *Index) Serve(ctx *faygo.Context) error {
	req := query.NewRequest()
	req.Connect(i.Mcsrvstatus)
	res, _ := req.Full()

	return ctx.JSON(200, res)
}

func main() {
	app := faygo.New("@maxoooz.dev/mcsrvstatus", "1.0.O")
	app.GET("/api/mcsrvstatus/:mcsrvstatus", new(Index))

	faygo.Run()
}