package main

import (
	"fmt"
	"time"
)

func greet(phrase string) {
	fmt.Println("Hello!", phrase)
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(5 * time.Second)
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func main() {
	//go greet("Nice to meet you")
	//go greet("How are you?")
	done := make(chan bool)
	go slowGreet("How ... are ... you ...?", done)
	//go greet("I hope you are well and safe!")
	<-done
}
