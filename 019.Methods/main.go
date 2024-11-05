package main

import "fmt"

type rect struct {
	width, height int
}

// 指针接收器
func (r *rect) area() int {
	return r.width * r.height
}

// 值接收器
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	//指向结构体的指针
	rp := &r
	fmt.Println("area: ", rp.area())
	//perim()方法是一个值接收者。当调用 rp.perim() 时，Go 会自动将 rp 解引用为其指向的结构体的值
	fmt.Println("perim:", rp.perim())
}
