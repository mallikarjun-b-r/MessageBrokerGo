package main

import (
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/gorilla/mux"
)

func TestHandler(t *testing.T) {
	//setup code for this test
	startDispatcher()
	//track time to end server after 3 secs
	now := time.Now()
	router := mux.NewRouter()
	router.HandleFunc("/payload", consumerHandler).Methods("POST")
	for i := 0; i < 100; i++ {
		messegeQueue <- "messege"
	}

	//func to close the server after 3s and send 100 messeges and verify if 100 messeges are recieved
	go func() {
		for {
			if time.Since(now).Seconds() >= 2 {
				if messegesProcessed == 100 {
					os.Exit(0)
				}
				os.Exit(1)
			}
		}
	}()

	//start http server
	log.Fatal(http.ListenAndServe(":8080", router))

}
