package main

import (
	"fmt"
)

func main() {
	decre(3)
	fmt.Println("after decre")
}
func decre(x int) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("internal error: %v\n", p)
		}
	}()
	fmt.Printf("decre(%d), res=%d\n", x+0/x, x)
	fmt.Printf("x=%d\n", x)
	decre(x - 1)
}
