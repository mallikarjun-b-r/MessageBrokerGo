package main

import (
	"testing"

	"github.com/valyala/fasthttp"
)

func TestRecieveMessege(t *testing.T) {
	ctx := fasthttp.RequestCtx{}
	ctx.Request.Header.Set("Content-Length", "7")
	ctx.Request.Header.Set("Content-Type", "application/json")

	fastHTTPHandlerGet(&ctx)
	if ctx.Response.StatusCode() != 200 {
		t.Fail()
	}
}

func TestRejectMessegeIfContentLengthIsMoreThan1Kb(t *testing.T) {
	ctx := fasthttp.RequestCtx{}
	ctx.Request.Header.Set("Content-Length", "1025")
	ctx.Request.Header.Set("Content-Type", "application/json")

	fastHTTPHandlerGet(&ctx)
	if ctx.Response.StatusCode() != 400 {
		t.Fail()
	}
}
func TestRejectMessegeIfMessegeQueueIsFull(t *testing.T) {

	ctx := fasthttp.RequestCtx{}
	for i := 0; i < 40001; i++ {
		messegeQueue <- "messege"
	}
	ctx.Request.Header.Set("Content-Length", "7")
	ctx.Request.Header.Set("Content-Type", "application/json")

	fastHTTPHandlerGet(&ctx)
	if ctx.Response.StatusCode() != 500 {
		t.Fail()
	}

}
