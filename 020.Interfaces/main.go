package main

import (
	"fmt"
	"math"
)

// 声明 一个 接口
type geometry interface {
	area() float64 // 方法
	perim() float64
}

// 声明 struct
type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// 为 struct 实现 geometry
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// 接收一个 实现了 geometry 的 结构体
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	measure(r) // 由于  rect 实现了  geometry 接口
	measure(c)
}
