package main

import (
	"time"
)

func send(c chan int) {
	c <- 1
	println("send,OK")
}

func revive(c chan int) {
	<-c
	println("revive,OK")
}

func main() {
	ch := make(chan int)
	ch <- 1
	go send(ch)
	go send(ch)

	go revive(ch)
	go revive(ch)
	go revive(ch)
	go revive(ch)

	time.Sleep(time.Hour)
}
