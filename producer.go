package main

import (
	"fmt"
	"io/ioutil"

	"github.com/buger/jsonparser"
	"github.com/cornelk/hashmap"
	"github.com/valyala/fasthttp"
)

var dir = "/tmp/big-o/file-%s"
var cache = &hashmap.HashMap{}

func fastHTTPHandlerPut(ctx *fasthttp.RequestCtx) {
	//os.Remove(fmt.Sprintf(dir, ctx.UserValue("probeId")))

	probeId := ctx.UserValue("probeId")
	body := ctx.PostBody()
	fileName := fmt.Sprintf("file-%s", probeId)
	transmissionTime, _ := jsonparser.GetInt(body, "eventTransmissionTime")
	if savedTime, ok := cache.Get(fileName); ok {
		var savedTransmissionTime = savedTime.(int64)
		// fmt.Print("saved ")
		// fmt.Println(savedTransmissionTime)
		// fmt.Print("current ")
		// fmt.Println(transmissionTime)

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

func writeToCache(fileName string, transmissionTime int64) {
	cache.Set(fileName, transmissionTime)
}
