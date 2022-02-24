package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"runtime"

	"github.com/buaazp/fasthttprouter"
	"github.com/buger/jsonparser"
	"github.com/valyala/fasthttp"
)

func main() {

	// Init Router
	router := fasthttprouter.New()
	//debug.SetGCPercent(-1)(disable gc if needed)
	files, err := ioutil.ReadDir("/data/")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		name := file.Name()
		data, err := ioutil.ReadFile("/data/" + name)
		if err != nil {
			panic(err)
		}
		transmissionTime, _ := jsonparser.GetInt(data, "eventTransmissionTime")
		// fmt.Println("file" + name)
		// fmt.Println(data)
		writeToCache(name, transmissionTime)
	}
	router.PUT("/probe/:probeId/event/:eventId", fastHTTPHandlerPut)
	router.GET("/probe/:probeId/latest", fastHTTPHandlerGet)

	fmt.Println("noOfGoRoutinesStart", runtime.NumGoroutine())
	log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}
