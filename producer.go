package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/buger/jsonparser"
	"github.com/cornelk/hashmap"
	"github.com/valyala/fasthttp"
)

var dir = "/data/file-%s"
var cache = &hashmap.HashMap{}

func fastHTTPHandlerPut(ctx *fasthttp.RequestCtx) {
	probeId := ctx.UserValue("probeId")
	body := ctx.PostBody()
	fileName := fmt.Sprintf("file-%s", probeId)
	transmissionTime, err := jsonparser.GetInt(body, "eventTransmissionTime")

	if (err != nil) {
		ctx.SetStatusCode(400)
		return
	}

	if savedTime, ok := cache.Get(fileName); ok {
		var savedTransmissionTime = savedTime.(int64)

		if savedTransmissionTime > transmissionTime {
			ctx.Response.SetStatusCode(200)
			return
		}
	}

	updated, err := jsonparser.Set(body, []byte(fmt.Sprintf("%d", time.Now().UnixMilli())), "eventReceivedTime")
	if (err != nil) {
		ctx.SetStatusCode(400)
		return
	}

	os.Remove(fmt.Sprintf(dir, ctx.UserValue("probeId")))
	ioutil.WriteFile(fmt.Sprintf(dir, ctx.UserValue("probeId")), updated, 0666)
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
