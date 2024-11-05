package main

import "fmt"

func intSeq() func() int { // 返回一个 函数,该函数返回int
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq() // 调用 intSeq，返回一个新的函数，并赋值给 nextInt
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	newInts := intSeq() // 再次调用 intSeq，返回另一个新的函数
	fmt.Println(newInts())
}
