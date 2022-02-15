package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func main() {

	// Init Router
	router := fasthttprouter.New()
	//debug.SetGCPercent(-1)(disable gc if needed)

	router.PUT("/probe/:probeId/event/:eventId", fastHTTPHandlerPost)
	router.GET("/probe/:probeId/latest", fastHTTPHandlerGet)

	fmt.Println("noOfGoRoutinesStart", runtime.NumGoroutine())
	log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}
