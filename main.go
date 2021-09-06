package main

import (
	"SE-Subscriber/usecases"
	"fmt"
	"log"
	"sync"

	"SE-Subscriber/repository"

	"github.com/nsqio/go-nsq"
)

func main() {

	es, err := repository.ESClient()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	fmt.Println("testing")
	config := nsq.NewConfig()
	q, _ := nsq.NewConsumer("test", "ch", config)

	q.AddHandler(&usecases.ReceiverMsg{Cli: es})
	err = q.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		log.Panic("Could not connect")
	}
	wg.Wait()
}
