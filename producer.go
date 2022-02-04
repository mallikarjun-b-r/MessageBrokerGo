package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"github.com/valyala/fasthttp"
)

func fastHTTPHandlerPost(ctx *fasthttp.RequestCtx) {
	message :=string(ctx.PostBody())
	num := rand.Intn(40000)
	ioutil.WriteFile(fmt.Sprintf("/Users/naman/Personalspace/temp/data/file-%d", num), 
	[]byte(fmt.Sprintf("%d\n%s", num, message)),0777)

	ctx.Response.SetBody([]byte(message))
	ctx.Response.SetStatusCode(200)
}

//Mainly to get the details about stats
func fastHTTPHandlerGet(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(200)
}
