package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("hello world")
	var wait sync.WaitGroup

	for i := 0; i < 4; i++ {
		wait.Add(1)
		go func(j int) {
			defer wait.Done()
			printSomething(j)
		}(i)
	}
	wait.Wait()
}
func printSomething(worker int) {
	fmt.Println(worker, " worker start....")
	time.Sleep(time.Second)
	fmt.Println(worker, " worker over....")

}
