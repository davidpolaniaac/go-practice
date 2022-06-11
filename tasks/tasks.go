package main

import "fmt"

type Task struct {
	name        string
	description string
	complete    bool
}

func (t *Task) setComplete() {
	t.complete = true
}

func (t *Task) setDescription(description string) {
	t.description = description
}

func (t *Task) setName(name string) {
	t.name = name
}
func main() {
	t := Task{
		name:        "test name",
		description: "test description",
	}

	t.setComplete()

	fmt.Printf("%+v\n", t)
}
