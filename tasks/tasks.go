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

type TaskList struct {
	tasks []*Task
}

func (tl *TaskList) addTask(t *Task) {
	tl.tasks = append(tl.tasks, t)
}

func (tl *TaskList) removeTask(index int) {
	tl.tasks = append(tl.tasks[:index], tl.tasks[index+1:]...)
}

func (tl *TaskList) print() {
	for i, v := range tl.tasks {
		fmt.Printf("index: %d - %+v\n", i, v)
	}
}

func (tl *TaskList) printComplete() {
	for i, v := range tl.tasks {
		if v.complete {
			fmt.Printf("index: %d - %+v\n", i, v)
		}
	}
}

func main() {
	t1 := &Task{
		name:        "name 1",
		description: "description 1",
	}
	t2 := &Task{
		name:        "name 2",
		description: "description 2",
	}

	t3 := &Task{
		name:        "name 3",
		description: "description 3",
	}

	listTasks := TaskList{
		tasks: []*Task{t1, t2},
	}

	listTasks.addTask(t3)

	fmt.Println(listTasks.tasks[0])

	listTasks.removeTask(0)
	listTasks.print()
	listTasks.tasks[0].complete = true
	listTasks.printComplete()

}
