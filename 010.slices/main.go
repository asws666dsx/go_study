package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string // 未初始化的切片等于 nil，长度为 0
	fmt.Println("unInit:", s, s == nil, len(s) == 0)

	//使用 make 函数创建一个长度为 3 的字符串切片，初始值为零值（空字符串）
	s = make([]string, 3)
	fmt.Println("slice:", len(s), cap(s))

	// 使用make 创建一个长度为3 ,容量为5的字符串切片
	s = make([]string, 3, 5)
	fmt.Println("slice", s, len(s), cap(s))

	// 可以像数组一样设置和get  slice
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	//和数组一样,len 会按预期返回切片的长度
	fmt.Println("len:", len(s))

	//一个是内置 append，它返回包含一个或多个新值的 slice
	s = append(s, "d")
	fmt.Println(s)
	s = append(s, "e", "f")
	fmt.Println(s)

	//切片也可以被复制。在这里，我们创建一个与 s 长度相同的空切片 c，并从 s 复制到 c 中
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	l := s[2:5] // 通过 下标获取到 s下标为2  到下标5的值
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l) // 通过下标获取到 s 下标 0 到下标 5的值

	l = s[2:]
	fmt.Println("sl3", l) // 通过 下标获取到 s 下标 2 到 最后一个的值

	// 也可以在一行中为 slice 声明和初始化一个变量
	t := []string{"a", "b", "c"}
	fmt.Println(t)

	t2 := []string{"a", "b", "c"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	//切片可以组合成多维数据结构。与多维数组不同，内部切片的长度可以变化
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d", twoD)
}
