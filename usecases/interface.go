package usecases

import (
	"SE-Subscriber/model"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/nsqio/go-nsq"
)

var Age int = 1

type ReceiverMsg struct {
	Cli *elasticsearch.Client
}

func (Rmsg *ReceiverMsg) HandleMessage(message *nsq.Message) error {
	fmt.Println("method being called during testing")
	var msg model.IMessage
	err := json.Unmarshal(message.Body, &msg)

	if err != nil {
		log.Fatal("error during marshalling: ", err)
	}
	if msg.Task == "created" {
		err = Rmsg.InsertInES(msg.Person)
		if err != nil {
			log.Fatal("error during insertion in ES", err)
		}
	}

	if msg.Task == "updated" {
		err = Rmsg.UpdateInES(msg.Person)
		if err != nil {
			log.Fatal("error during insertion in ES", err)
		}
	}

	if msg.Task == "deleted" {
		err = Rmsg.DeleteInES(msg.Person.DocumentId)
		if err != nil {
			log.Fatal("error during insertion in ES", err)
		}
	}
	return nil
}

func (Rmsg *ReceiverMsg) InsertInES(person model.Person) error {

	dataJson, err := json.Marshal(person)
	if err != nil {
		log.Fatal("error during marshalling: ", err)
	}

	stringMess := string(dataJson)
	log.Println("string during: ", stringMess)
	request := esapi.IndexRequest{Index: "person", DocumentID: strconv.Itoa(person.DocumentId), Body: strings.NewReader(stringMess)}

	res, err := request.Do(context.Background(), Rmsg.Cli)
	if err != nil {
		log.Fatal("error during insertion in ES", err)
	}
	log.Println("response of es", res)
	return nil
}

func (Rmsg *ReceiverMsg) UpdateInES(person model.Person) error {

	dataJson, err := json.Marshal(person)
	if err != nil {
		log.Fatal("error during marshalling: ", err)
	}

	stringMess := string(dataJson)
	log.Println("string during: ", stringMess)
	request := esapi.IndexRequest{Index: "person", DocumentID: strconv.Itoa(person.DocumentId), Body: strings.NewReader(stringMess)}

	res, err := request.Do(context.Background(), Rmsg.Cli)
	if err != nil {
		log.Fatal("error during insertion in ES", err)
	}
	log.Println("response of es", res)
	return nil
}

func (Rmsg *ReceiverMsg) DeleteInES(Id int) error {

	request := esapi.DeleteRequest{Index: "person", DocumentID: strconv.Itoa(Id)}
	res, err := request.Do(context.Background(), Rmsg.Cli)
	if err != nil {
		log.Fatal("error during insertion in ES", err)
	}
	log.Println("response of es", res)
	return nil
}
