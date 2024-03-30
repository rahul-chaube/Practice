package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)
	ch := make(chan bool)
	go PrintOdd(ch)
	go PrintEven(ch)

	wg.Wait()

	// ========  Using Two channel
	wg.Add(2)
	no_chan := make(chan int)
	is_done := make(chan bool)

	go printEvent(no_chan, is_done)
	go printOdd(no_chan, is_done)

	for i := 1; i <= 10; i++ {
		no_chan <- i
		<-is_done
	}
	close(no_chan)
	close(is_done)
	wg.Wait()

}
func printEvent(no_chan chan int, is_done chan bool) {
	defer wg.Done()

	for no := range no_chan {
		if no%2 == 0 {
			fmt.Println("event number ", no)
			is_done <- true
		} else {
			no_chan <- no
		}
	}

}
func printOdd(no_chan chan int, is_done chan bool) {
	defer wg.Done()

	for no := range no_chan {
		if no%2 != 0 {
			fmt.Println("Odd number ", no)
			is_done <- true
		} else {
			no_chan <- no
		}
	}

}

func PrintOdd(ch chan bool) {
	defer wg.Done()
	for i := 1; i <= 10; i += 1 {
		ch <- true
		if i%2 == 1 {
			fmt.Println("Odd Number ", i)
		}
	}
}
func PrintEven(ch chan bool) {
	defer wg.Done()
	for i := 1; i <= 10; i += 1 {
		<-ch
		if i%2 == 0 {
			fmt.Println("Even Number ", i)
		}
	}
}
