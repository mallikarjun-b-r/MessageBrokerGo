package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	"github.com/valyala/fasthttp"
)

func fastHTTPHandlerPost(ctx *fasthttp.RequestCtx) {
	message := string(ctx.PostBody())
	num := rand.Intn(100000)
	dt := time.Now().String()[18:19]

	file_name := fmt.Sprintf("/Users/naman/Personalspace/temp/data/file-%d-%s", num, dt)

	os.Remove(file_name)
	ioutil.WriteFile(file_name, []byte(fmt.Sprintf("%d\n%s", num, message)), 0777)

	// ctx.Response.SetBody([]byte(message))
	ctx.Response.SetStatusCode(200)
}

//Mainly to get the details about stats
func fastHTTPHandlerGet(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(200)
}
