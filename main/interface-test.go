package main

import (
	"fmt"
	"io"
	"os"
)

// 方法和函数区别，方法有一个调度器，也即对象

type ByteCounter int

// 该方法实现io.Writer接口中的Writer方法，所以打印出自定义的字节长度
func (b *ByteCounter) Write(p []byte) (int, error) {
	*b += ByteCounter(len(p))
	return 10, nil
}
func main() {
	var w io.Writer
	w = os.Stdout //  os.Stdout为io.Writer接口实现

	w.Write([]byte("eeeee\n"))
	//fmt.Printf("w= %v \n",w)
	var c ByteCounter
	c.Write([]byte("hello!!"))
	fmt.Printf("c= %v \n", c)
}
