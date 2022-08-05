package main

import "fmt"

type Topic interface {
	registry(observer Observer)
	broadcast()
}

type Observer interface {
	getId() string
	updateValue(value string)
}

type Item struct {
	observers []Observer
	name      string
	available bool
}

func (i *Item) register(observer Observer) {
	i.observers = append(i.observers, observer)
}

func (i *Item) broadcast() {
	for _, observer := range i.observers {
		observer.updateValue(i.name)
	}
}

func NewItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (i *Item) UpdateAvailable() {
	fmt.Printf("Item %s is available\n", i.name)
	i.available = true
	i.broadcast()
}

type EmailClient struct {
	id string
}

func (eC *EmailClient) updateValue(value string) {
	fmt.Printf("Sending Email - %s available from client %s\n", value, eC.id)
}

func (eC *EmailClient) getId() string {
	return eC.id
}

func main() {
	nvidiaItem := NewItem("RTX 3080")
	firstObserver := &EmailClient{
		id: "1234",
	}
	secondObserver := &EmailClient{
		id: "56789",
	}
	nvidiaItem.register(firstObserver)
	nvidiaItem.register(secondObserver)
	nvidiaItem.UpdateAvailable()
}
