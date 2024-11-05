package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	//像 & & 和 || 这样的逻辑运算符在条件表达式中通常很有用
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 of are even")
	}

	//可以在位于条件表达式之前声明变量.此语句中声明的任何变量在 Current 和所有后续分支中都可用
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
