package main

import "fmt"

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 2)
	pongs := make(chan string, 2)
	go pong(pings, pongs)
	pings <- "hello"
	pings <- " world"
	fmt.Print(<-pongs)
}
