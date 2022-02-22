package main

import (
	"fmt"
	"io/ioutil"

	"github.com/buger/jsonparser"
	"github.com/valyala/fasthttp"
)

var dir = "/tmp/big-o/file-%s"
var cache = make(map[string]string)

func fastHTTPHandlerPut(ctx *fasthttp.RequestCtx) {
	//os.Remove(fmt.Sprintf(dir, ctx.UserValue("probeId")))
	probeId := ctx.UserValue("probeId")
	body := ctx.PostBody()
	fileName := fmt.Sprintf("file-%s", probeId)
	transmissionTime, _ := jsonparser.GetString(body, "eventTransmissionTime")
	if savedTransmissionTime, ok := cache[fileName]; ok {
		fmt.Println("saved" + savedTransmissionTime)
		fmt.Println("current" + transmissionTime)

		if savedTransmissionTime > transmissionTime {
			ctx.Response.SetStatusCode(200)
			return
		}
	}

	ioutil.WriteFile(fmt.Sprintf(dir, ctx.UserValue("probeId")), ctx.PostBody(), 0666)
	writeToCache(fileName, transmissionTime)
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

func writeToCache(fileName string, transmissionTime string) {
	cache[fileName] = transmissionTime
}
