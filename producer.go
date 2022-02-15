package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/valyala/fasthttp"
)

var dir = "/Users/naman/Personalspace/temp/data/file-%s"

func fastHTTPHandlerPost(ctx *fasthttp.RequestCtx) {
	os.Remove(fmt.Sprintf(dir, ctx.UserValue("probeId")))
	ioutil.WriteFile(fmt.Sprintf(dir, ctx.UserValue("probeId")), ctx.PostBody(), 0666)
	ctx.Response.SetStatusCode(200)
}

func fastHTTPHandlerGet(ctx *fasthttp.RequestCtx) {
	file, err := ioutil.ReadFile(fmt.Sprintf(dir, ctx.UserValue("probeId")))
	if err != nil {
		ctx.SetStatusCode(404)
	} else {
		ctx.Response.SetBody(file)
		ctx.SetStatusCode(200)
	}
}
