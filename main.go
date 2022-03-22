package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/buaazp/fasthttprouter"
	"github.com/buger/jsonparser"
	"github.com/valyala/fasthttp"
)

func main() {

	// Init Router
	router := fasthttprouter.New()
	//debug.SetGCPercent(-1)(disable gc if needed)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		go loadFile(file.Name())
	}
	router.PUT("/probe/:probeId/event/:eventId", fastHTTPHandlerPut)
	router.GET("/probe/:probeId/latest", fastHTTPHandlerGet)

	fmt.Println("noOfGoRoutinesStart", runtime.NumGoroutine())
	log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}

func loadFile(fileName string) {
	if(strings.HasSuffix(fileName, ".swp")) {
		nameWithoutSwp := strings.Replace(fileName, ".swp", "", 1)
		os.Remove(dir + nameWithoutSwp)
		os.Rename(dir + fileName, dir + nameWithoutSwp)
		fileName = nameWithoutSwp
	}

	data, err := ioutil.ReadFile(dir + fileName)
	if err != nil {
		panic(err)
	}

	transmissionTime, _ := jsonparser.GetInt(data, "eventTransmissionTime")
	writeToCache(fileName, transmissionTime)
}