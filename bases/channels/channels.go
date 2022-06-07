package channels

import "fmt"

func say(text string, c chan<- string) {
	c <- text
}

func Run() {
	c := make(chan string, 1)

	fmt.Println("Hello")

	go say("Bye", c)

	fmt.Println(<-c)
}

func message(text string, c chan string) {
	c <- text
}

func Multi() {
	c := make(chan string, 2)

	c <- "message 1"
	c <- "message 2"

	fmt.Println(len(c), cap(c))

	// Range y close
	close(c)
	//c<-"messagecd 3"

	for message := range c {
		fmt.Println(message)
	}

	// Select
	email1 := make(chan string)
	email2 := make(chan string)

	go message("message 1", email1)
	go message("message 2", email2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email from channel 1", m1)
		case m2 := <-email2:
			fmt.Println("Email from channel 2", m2)
		}
	}
}
