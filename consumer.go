package main

import (
	"log"
)

var consumerSlice []string
var strPost = []byte("POST")

// Consumer struct
type Consumer struct {
	fqdn          string
	messeges      chan string
	consumerQueue chan chan string
	quitWorker    chan bool
}

// NewConsumer creates, and return the consumer.
func NewConsumer(fqdn string, consumerQueue chan chan string) Consumer {
	consumer := Consumer{
		fqdn:          fqdn,
		messeges:      make(chan string),
		consumerQueue: consumerQueue,
		quitWorker:    make(chan bool)}

	return consumer
}

//Starts the consumer and adds its channel to consumerQueue to recieve messeges
func (c *Consumer) start() {
	go func() {
		for {
			//Adding the current consumer.messeges to consumerQueue to accept messeges from dispatcher
			c.consumerQueue <- c.messeges
			select {
			case messege := <-c.messeges:
				sendMessege(messege, c)
			case <-c.quitWorker:
				return
			}
		}
	}()
}

func sendMessege(messege string, c *Consumer) {
}

func (c *Consumer) stop() {
	log.Println("Stopping Consumer ", c.fqdn)
}
