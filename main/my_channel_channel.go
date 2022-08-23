package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("hello world")
	ch := make(chan string, 20)
	num := 25
	sendDoneChannel := make(chan bool, num)
	recvDoneChannel := make(chan bool, num)

	for i := 0; i < num; i++ {
		go producer(i, ch, sendDoneChannel)
	}
	for i := 0; i < num; i++ {
		go consumer(i, ch, recvDoneChannel)
	}
	for i := 0; i < num; i++ {
		<-sendDoneChannel
	}
	close(sendDoneChannel)
	close(ch)
	for i := 0; i < num; i++ {
		<-recvDoneChannel
	}
	close(recvDoneChannel)
}
func producer(producerId int, ch chan string, sendDoneChannel chan bool) {
	for i := 0; i < 10; i++ {
		str := strconv.Itoa(i)
		ch <- str
		fmt.Printf(" %d product send %s \n", producerId, str)
	}
	sendDoneChannel <- true
}
func consumer(consumerId int, ch chan string, recvDoneChannel chan bool) {
	for i := range ch {
		fmt.Printf(" %d consumer receive %s \n", consumerId, i)
	}
	recvDoneChannel <- true
}
