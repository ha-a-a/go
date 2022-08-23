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
	var w io.Writer //接口的零值 = 动态类型和动态值均为nil
	if w == nil {
		fmt.Printf("w = nil \n")
	}
	w = os.Stdout //  os.Stdout为io.Writer接口实现
	w.Write([]byte("eeeee\n"))
	//fmt.Printf("w= %v \n",w)
	var c ByteCounter
	c.Write([]byte("hello!!"))
	fmt.Printf("c= %v \n", c)
	// 类型断言,具体类型断言,断言返回值为w的动态类型*os.File
	f := w.(*os.File)
	fmt.Printf("f=%T \n", f)
	// 类型断言，接口类型断言,断言返回值为w的动态类型*os.File，并且拥有接口io.ReadWriter的动态类型和动态值
	var iw io.Writer
	iw = os.Stdout
	rw := iw.(io.ReadWriter)
	fmt.Printf("rw=%T \n", rw)

}
