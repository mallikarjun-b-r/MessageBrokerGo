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

var dir = "/tmp/data/"
var fileTemplate = dir + "file-%s"
var cache = &hashmap.HashMap{}

func fastHTTPHandlerPut(ctx *fasthttp.RequestCtx) {
	probeId := ctx.UserValue("probeId")
	body := ctx.PostBody()
	fileName := fmt.Sprintf("file-%s", probeId)
	transmissionTime, err := jsonparser.GetInt(body, "eventTransmissionTime")

	if err != nil {
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
	if err != nil {
		ctx.SetStatusCode(400)
		return
	}

	os.Remove(fmt.Sprintf(fileTemplate, ctx.UserValue("probeId")))
	ioutil.WriteFile(fmt.Sprintf(fileTemplate, ctx.UserValue("probeId")), updated, 0666)
	writeToCache(fileName, transmissionTime)
}

func fastHTTPHandlerGet(ctx *fasthttp.RequestCtx) {
	file, err := ioutil.ReadFile(fmt.Sprintf(fileTemplate, ctx.UserValue("probeId")))
	if err != nil {
		ctx.SetStatusCode(404)
	}
	
	ctx.Response.SetBody(file)
}

func writeToCache(fileName string, transmissionTime int64) {
	cache.Set(fileName, transmissionTime)
}
