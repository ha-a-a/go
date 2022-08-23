package main

import (
	"fmt"
	"strconv"
	"sync"
)

var channel = make(chan string, 20)
var wait sync.WaitGroup
var waitRecv sync.WaitGroup

func main() {
	fmt.Println("hello world")
	for i := 0; i < 25; i++ {
		wait.Add(1)
		go send(i)
	}
	for i := 0; i < 25; i++ {
		waitRecv.Add(1)
		go recv(i)
	}
	wait.Wait()
	close(channel)
	waitRecv.Wait()
}
func send(producer int) {
	for i := 0; i < 10; i++ {
		str := strconv.Itoa(i)
		channel <- str
		fmt.Printf(" %d product send %s \n", producer, str)
	}
	wait.Done()
}
func recv(consumer int) {
	for i := range channel {
		fmt.Printf(" %d consumer receive %s \n", consumer, i)
	}
	waitRecv.Done()
}
