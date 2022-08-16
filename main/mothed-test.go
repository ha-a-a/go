package main

import "fmt"

// 方法和函数区别，方法有一个调度器，也即对象

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
	var po func(p, offset Point) Point
	if add {
		po = Point.Add
	} else {
		po = Point.Sub
	}
	for i := range p {
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
