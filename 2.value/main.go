package main

import "fmt"

func main() {
	// 字符串可以通过 + 将两个字符串进行拼接
	fmt.Println("go" + "lang")
	// int类型的加减法
	fmt.Println("1+1=", 1+1)
	// 浮点型的除法
	fmt.Println("7.0/3.0=", 7.0/3.0)

	// true 就是正确的意思 而 false 为 错误的意思
	// &&为与 表示：条件中都为true 的情况下,整个表达式才为true 否则为 false
	fmt.Println(true && false)
	// ||为或  表示  条件中只要有一个为true,整个表达式就为true
	fmt.Println(true || false)
	// ！为非   用于取反布尔值
	fmt.Println(!true)

}
