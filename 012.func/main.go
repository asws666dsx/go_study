package main

import "fmt"

// 函数和 变量的 首字母小写表示只能在当前包中使用
// 这是一个函数，它接受两个 int 并将它们的总和作为 int 返回
func plus(a int, b int) int {
	return a + b
}

// Go 需要显式返回，即它不会自动返回最后一个表达式的值
func plusPlus(a, b, c int) int { // 当有多个相同类型的连续参数时，可以省略类似类型参数的类型名称，直到声明该类型的最终参数。
	return a + b + c
}

// 此函数签名中的 （int， int） 显示该函数返回 2 ints
func vals() (int, int) {
	return 3, 7
}

// 可变参数函数
func sums(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	res := plus(1, 2) //调用方法
	fmt.Println("1+2=", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3=", res)

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals() //如果只需要返回值的子集，请使用空标识符 _
	fmt.Println(c)

	sums(1, 2)
	sums(1, 2, 3)
	nums := []int{1, 2, 3, 4, 5, 6}
	sums(nums...) // nums... 将切片展开，使每个元素作为独立的参数传递给 sums
}
