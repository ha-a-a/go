package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// select case类似switch case，不同点：select 会等待有可执行的case，当多个case都能执行时，会随机一个case执行,
	//default无可执行的case时，走default，实现select无阻塞操作
	//launchOrAbort()
	//randomCase()
	selectDefault()
}
func selectDefault() {
	ch := make(chan struct{})
	go func() {
		fmt.Println("waiting for input....")
		os.Stdin.Read(make([]byte, 1))
		ch <- struct{}{}
	}()
	select {
	case <-ch:
		fmt.Println("input something")
		return
	default:
		fmt.Println("waiting for input something...")
	}
}
func randomCase() {
	ch := make(chan int, 2)
	for i := 0; i < 50; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}

	}
}
func launchOrAbort() {
	abort := make(chan struct{})
	go func() {
		fmt.Println("waiting for abort....")
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	select {
	case <-time.After(10 * time.Second):
		fmt.Println("launch successful!")
	case <-abort:
		fmt.Println("Lanch aborted!")

	}
}
