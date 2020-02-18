package main

import (
	"bufio"
	"log"
	"os"
	"runtime"
)

var consumerQueue chan chan string

//Starts the dispating of messeges by creating 10 concurrent consumers for each url
func startDispatcher() {
	loadConsumersFromFile()
	noOfConsumers := len(consumerSlice)

	//Create Consumer Queue
	consumerQueue = make(chan chan string, 50*noOfConsumers)

	//Create All Consumers
	for i := 0; i < 50*noOfConsumers; i++ {
		log.Println("Starting Consumer ", i, " ", consumerSlice[i%noOfConsumers])
		consumer := NewConsumer(consumerSlice[i%noOfConsumers], consumerQueue)
		//Starting the Consumer
		consumer.start()
	}

	//If a messges is available in the buffer channel sends it to a first available conusmer to process it
	go func() {
		for {
			if runtime.NumGoroutine() > 300000 {
				continue
			}
			select {
			case messege := <-messegeQueue:
				//log.Println("Recieved Messege ", messege)
				go func() {
					consumer := <-consumerQueue
					//log.Println("Dispatching Messege")
					//fmt.Println("noOfGoRoutinesStart", runtime.NumGoroutine())
					consumer <- messege
				}()
			}
		}
	}()
}

//Loads the consumerUrl into ConsumerSlice from a file
//Change the path to where the consumers.txt file is
func loadConsumersFromFile() {
	file, err := os.Open("/Users/mallikarjunbr/golang/broker/consumers.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		consumerSlice = append(consumerSlice, scanner.Text())
		//fmt.Println(consumerSlice[i])
		i++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
