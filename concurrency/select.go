package concurrency

import (
	"fmt"
	"time"
)

var hellos = []string{"Hello!", "Namaste", "Ciao"}
var goodbyes = []string{"Goodbye!", "Alvida!", "Arrivederci!"}

func ChannelSelect() {

	ch := make(chan string, 1)
	ch2 := make(chan string, 1)

	go greetSelect(hellos, ch)
	go greetSelect(goodbyes, ch2)

	time.Sleep(1 * time.Second)
	fmt.Println("Main ready:")

	for {
		select {
		case gr, ok := <-ch:
			if !ok {
				ch = nil
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

func greetSelect(greetings []string, ch chan<- string) { // make the channel send-only for this function by putting <- after the chan keyword. No symbol means bi-directional
	fmt.Printf("Greeter ready! \n Greeter ready to send greeting...\n")

	for _, g := range greetings {
		ch <- g
	}
	close(ch)

	fmt.Println("Greeter completed")
}

func printGreeting(greeting string) {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Greeting received!", greeting)
}
