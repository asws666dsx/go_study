package main

import "fmt"

func zeroVal(iVal int) {
	iVal = 0
}

// *int 为接收远古 int类型的变量的地址
func zeroPtr(iPtr *int) {
	*iPtr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)
	zeroVal(i) // 由于是值拷贝所以原来的i没变
	fmt.Println("zeroVal:", i)
	zeroPtr(&i) // & 为 传递地址,因为传递的是地址,所有这个函数里修改的是i 本身
	fmt.Println("zeroPtr:", i)
	fmt.Println("zeroPtr:", &i)
}
