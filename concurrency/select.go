package concurrency

import (
	"fmt"
	"time"
)

var hellos = []string{"Hello", "Ciao", "Salut", "Namaste"}
var goodbyes = []string{"Goodbye", "Arrivederci", "Adios", "Namastey"}

func SelectFunc() {

	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	go greets(hellos, ch1)
	go greets(goodbyes, ch2)

	time.Sleep(1 * time.Second)
	fmt.Println("Main ready:")

	for {
		select {
		case gr, ok := <-ch1:
			// to avoid infinite loop, sets the channel to nil if it is closed.
			if !ok {
				ch1 = nil
				break
			}
			printGreeting(gr)
		case gr2, ok := <-ch2:
			if !ok {
				ch2 = nil
				break
			}
			printGreeting(gr2)
		default:
			return
		}

	}

}

func greets(greetings []string, ch chan<- string) {
	fmt.Println("Greeter ready")
	for _, g := range greetings {
		ch <- g
	}
	close(ch) // close channel to avoid goroutine locks
	fmt.Println("Greeter completed")
}

func printGreeting(greeting string) {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Greeting received:", greeting)
}
