package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	no_chan := make(chan int)
	is_done := make(chan bool)

	go even(no_chan, is_done)
	go odd(no_chan, is_done)

	for i := 0; i <= 10; i++ {
		no_chan <- i
		<-is_done

	}
	close(is_done)
	close(no_chan)
	fmt.Println("Exit ")

	fmt.Println("=======================  Save =========================== ")

	ch := make(chan struct{})
	wg.Add(2)

	go Even(ch)
	go Odd(ch)

	wg.Wait()

}

func even(no_chan chan int, isDone chan bool) {
	for data := range no_chan {
		if data%2 == 0 {
			fmt.Println("Number is even ", data)
			isDone <- true
		} else {
			no_chan <- data
		}
	}
}

func odd(no_chan chan int, isDone chan bool) {
	for data := range no_chan {
		if data%2 != 0 {
			fmt.Println("Number is odd ", data)
			isDone <- true
		} else {
			no_chan <- data
		}
	}
}

func Even(ch chan struct{}) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch <- struct{}{}
		if i%2 == 0 {
			fmt.Println("Event number is ", i)
		}
	}

}

func Odd(ch chan struct{}) {
	wg.Done()
	for i := 0; i < 10; i++ {
		<-ch
		if i%2 != 0 {
			fmt.Println("Odd number is ", i)
		}
	}
}
