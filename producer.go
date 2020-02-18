package main

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/valyala/fasthttp"
)

var messegeQueue = make(chan string, 60000)
var prevc = 0
var prevt int64 = 0
var counter int64 = 0
var sum int64 = 0
var rateSlice []string

func fastHTTPHandlerPost(ctx *fasthttp.RequestCtx) {
	contentLen, _ := strconv.ParseInt(string(ctx.Request.Header.Peek("Content-Length")), 10, 64)

	if contentLen > 1024 {
		ctx.Error("Messege legth exceeds 1024 chars rejecting messege ", 400)
		return
	}
	log.Println(len(messegeQueue))
	if len(messegeQueue) > 40000 {
		ctx.Error("Consumers are slow broker will not accept messeges", 500)
		return
	}

	totalMessegesRcvid++
	messege := string(ctx.PostBody())
	messegeQueue <- messege
	ctx.Response.SetStatusCode(200)
	//fmt.Fprintf(ctx., "Hi there! RequestURI is %q", ctx.RequestURI())
}

//Mainly to get the details about stats
func fastHTTPHandlerGet(ctx *fasthttp.RequestCtx) {
	//length := strconv.Itoa(messegesProcessed)

	tim := time.Now().UnixNano()
	//timeNano := strconv.FormatInt(tim, 10)
	//totalm := strconv.Itoa(totalMessegesRcvid)
	rate := (int64(messegesProcessed-prevc) * (1000000000) / (tim - prevt))
	//val1 := (tim - prevt)
	//val := messegesProcessed - prevc

	prevc = messegesProcessed
	prevt = tim

	rateSlice = append(rateSlice, strconv.FormatInt(rate, 10))
	counter++
	sum += rate

	ctx.SetStatusCode(200)
	ctx.Response.SetBody([]byte("messegeRate" + " " + strconv.FormatInt(sum/counter, 10) + " Failed Messeges " + strconv.Itoa(failedMesseges) + " Rates History " + strings.Join(rateSlice, ",")))
}
