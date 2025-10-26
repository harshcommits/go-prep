package concurrency

import (
	"fmt"
	"time"
)

var greetings = []string{"Hello!", "Namaste", "Ciao"}

func Channel() {
	ch := make(chan string, 1)
	go greet(ch)

	time.Sleep(5 * time.Second)
	fmt.Println("Main ready:")

	// not-so-efficient method
	// for {
	// 	greeting, ok := <-ch
	// 	if !ok {
	// 		return // this is to stop the flow if the channel is empty; !ok means channel is closed now
	// 	}
	// 	time.Sleep(2 * time.Second)
	// 	fmt.Println("Greeting received!", greeting)
	// }

	for greeting := range ch {
		fmt.Println("Greeting received: ", greeting)
	}
}

func greet(ch chan<- string) { // make the channel send-only for this function by putting <- after the chan keyword. No symbol means bi-directional
	fmt.Printf("Greeter ready! \n Greeter ready to send greeting...\n")

	for _, g := range greetings {
		ch <- g
	}
	close(ch)

	fmt.Println("Greeter completed")
}
