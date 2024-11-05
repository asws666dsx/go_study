package main

import "fmt"

func main() {
	//创建一个数组 a，它正好容纳 5 个 int。元素的类型和长度都是数组类型的一部分。默认情况下，数组的值为零，
	var a [5]int
	fmt.Println("emp", a)

	// 可以使用array[index] = value Y语法在索引处设置一个值，并使用array[index] 获取一个值
	a[4] = 100
	fmt.Println("Set", a)
	fmt.Println("get", a[4])

	// 内置的len 返回数组的长度
	fmt.Println("len:", len(a))

	//使用此语法可在一行中声明和初始化数组
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	//也可以让编译器为你计算元素的数量 ...
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// 如果使用 ： 指定索引，则两者之间的元素将归零,这里指定 第三个为400
	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx", b)
}
