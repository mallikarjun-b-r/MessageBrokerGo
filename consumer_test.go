package main

import (
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/gorilla/mux"
)

func TestStartConsumer(t *testing.T) {
	loadConsumersFromFile()
	consumerQueue = make(chan chan string, 1*10)
	consumer := NewConsumer(consumerSlice[0], consumerQueue)
	consumer.start()
	consumer.messeges <- "messege"
	consumer.stop()
	if len(consumerQueue) == 0 {
		t.Fail()
	}
}
func TestSendMessege(t *testing.T) {
	//setup code for this test
	loadConsumersFromFile()
	consumerQueue = make(chan chan string, 1*10)
	consumer := NewConsumer("http://localhost:8080", consumerQueue)

	//track time to end server after 3 secs
	now := time.Now()

	router := mux.NewRouter()
	router.HandleFunc("/payload", consumerHandler).Methods("POST")
	//func to close the server after 3s and send 100 messeges and verify if 100 messeges are recieved
	go func() {
		for {
			if time.Since(now).Seconds() >= 1 {
				for i := 0; i < 100; i++ {
					sendMessege("messege", &consumer)
				}
				if messegesProcessed == 101 || messegesProcessed == 100 {
					os.Exit(0)
				}
				os.Exit(1)
			}
		}
	}()

	//start http server dummy consumer
	log.Fatal(http.ListenAndServe(":8080", router))

}

//dummyHandler
func consumerHandler(w http.ResponseWriter, r *http.Request) {
}
