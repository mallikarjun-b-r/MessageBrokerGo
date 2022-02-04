package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

var messegesRecievedAtConsumer = 0

var messegesProcessed = 0
var totalMessegesRcvid = 0
var failedMesseges = 0
var wg sync.WaitGroup

func main() {

	// Init Router
	router := fasthttprouter.New()
	//debug.SetGCPercent(-1)(disable gc if needed)

	router.POST("/payload", fastHTTPHandlerPost)
	router.GET("/stats", fastHTTPHandlerGet)

	//trace.Start(f)
	//defer trace.Stop()
	startDispatcher()

	fmt.Println("noOfGoRoutinesStart", runtime.NumGoroutine())
	log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}
