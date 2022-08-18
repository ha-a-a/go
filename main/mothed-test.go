package main

import "fmt"

// 方法和函数区别，方法有一个调度器，也即对象
//1）方法值：隐式调用, struct实例获取方法对象
//2)  方法表达式：显示调用, struct类型获取方法对象, 需要传递struct实例对象作为参数。
type Point struct {
	x, y float64
}

func (p Point) Add(q Point) Point {
	return Point{p.x + q.x, p.y + q.y}
}
func (p Point) Sub(q Point) Point {
	return Point{p.x - q.x, p.y - q.y}
}

type Path []Point

func (p Path) TranslateBy(offset Point, add bool) {
	// po为func(p, offset Point) Point的函数值
	var po func(p, offset Point) Point
	if add {
		// Point类型的Add方法值
		po = Point.Add
	} else {
		po = Point.Sub
	}
	for i := range p {
		// Point类型的Add或Sub方法表达式
		p[i] = po(p[i], offset)
	}
}
func main() {
	path := Path{Point{1, 2}, Point{2, 3},
		Point{4, 5}}
	fmt.Printf("path=%v \n", path)
	path.TranslateBy(Point{1, 1}, true)
	fmt.Printf("after add path=%v \n", path)
	path.TranslateBy(Point{1, 1}, false)
	fmt.Printf("after sub path=%v \n", path)
}
