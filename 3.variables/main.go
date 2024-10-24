package main

import "fmt"

func main() {
	// var 可以用来声明一个或多个变量
	var a = "initial" //让编译器自己去推导数据类型
	fmt.Println(a)

	var b, c int = 1, 2 // 指定类型
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int      // 声明一个变量,不赋予值.
	fmt.Println(e) // 输出e 这个变量的默认值.

	f := "apple" // := 语法是声明和初始化变量的简写, f:= "apple" 与  var f string = "apple" 表示的意思是一样的
	fmt.Println(f)
}
