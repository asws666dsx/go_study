package main

import "fmt"

func main() {
	// 自动在输出的尾部添加换行符\n,并自动处理参数类型
	fmt.Println("1.hello world")

	/*
		不会自动添加换行符,除非你手动在格式字符串中指定
		fmt.Printf("%s", "1.hello world")
	*/
}
