package model

type Person struct {
	Name       string
	Age        int
	Email      string
	DocumentId int
}

type IMessage struct {
	Task   string
	Person Person
}
