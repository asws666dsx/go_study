package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	//无类型常量可以在需要特定类型的上下文中自动转换为适当的类型
	const n = 500000000 // 声明一个无类型整数

	// e 为科学计数法中指数的表示,在科学计数法中,e 表示"10"的幂 这里表示 3 × 10^20
	const d = 3e20 / n

	fmt.Println(int64(d))

	// math.Sin 用于计算角度的正弦值. 接收一个浮点数.go 编译器将常量n 转为了浮点型
	fmt.Println(math.Sin(n))
}
