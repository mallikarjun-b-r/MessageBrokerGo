package main

import (
	"fmt"
	"io/ioutil"

	"github.com/valyala/fasthttp"
)

var dir = "/Users/naman/Personalspace/temp/data/file-%s"

func fastHTTPHandlerPost(ctx *fasthttp.RequestCtx) {
	ioutil.WriteFile(fmt.Sprintf(dir, ctx.UserValue("probeId")), ctx.PostBody(), 0666)
	ctx.Response.SetStatusCode(200)
}

//Mainly to get the details about stats
func fastHTTPHandlerGet(ctx *fasthttp.RequestCtx) {
	file, err := ioutil.ReadFile(fmt.Sprintf(dir, ctx.UserValue("probeId")))
	if err != nil {
		ctx.SetStatusCode(404)
	} else {
		ctx.Response.SetBody(file)
		ctx.SetStatusCode(200)
	}
}
