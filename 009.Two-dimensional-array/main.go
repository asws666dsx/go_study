package main

import "fmt"

func main() {
	var twoD [2][3]int //创建一个二维数组
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d", twoD)

	//也可以一次创建和初始化多维数组
	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d", twoD)
}
