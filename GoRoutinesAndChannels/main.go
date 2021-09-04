package main

import (
	"fmt"
	"time"
)

func fibonacci_select(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Quit.")
			return
		}
	}
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")

	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)

	bufferedChannel := make(chan int, 2)
	bufferedChannel <- 1
	bufferedChannel <- 2
	// bufferedChannel <- 3 THIS OVERFLOWS THE BUFFER AND CAUSES A DEADLOCK!
	fmt.Println(<-bufferedChannel)
	fmt.Println(<-bufferedChannel)

	fibonacciChannel := make(chan int, 10)
	go fibonacci(cap(fibonacciChannel), fibonacciChannel)
	for i := range fibonacciChannel {
		fmt.Println(i)
	}
	res, isOpen := <-fibonacciChannel
	fmt.Println(res, isOpen)

	selectChannel := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-selectChannel)
		}
		quit <- 0
	}()
	fibonacci_select(selectChannel, quit)

	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("     .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
