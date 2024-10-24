package main

import "fmt"

func main() {
	i := 1

	// 写法一
	for i < 3 { // 当条件满足时进入循环，直到条件不满足
		fmt.Println(i)
		i = i + 1 // 每次循环后 i 增加 1
	}

	// 写法二
	for j := 0; j < 3; j++ { // 从 0 到 2（包括 0，不包括 3）的循环
		fmt.Println(j)
	}

	// 写法三
	for i := range 3 {
		fmt.Println("range", i) // 输出索引值
	}

	// 写法四
	for { // 无限循环
		fmt.Println("loop")
		break // 立即跳出循环
	}

	for n := 0; n < 6; n++ { // 从 0 到 5 的循环
		if n%2 == 0 { // 当 n 为偶数时
			continue // 跳过本次循环，继续下一个循环
		}
		fmt.Println(n) // 输出奇数
	}
}
