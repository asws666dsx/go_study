package main

import (
	"fmt"
	"time"
)

func main() {

	// 基本用法
	i := 2
	fmt.Println("Write", i, "as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("tow")
	case 3:
		fmt.Println("three")
	}

	// 使用逗号分隔同一case 语句中的多个表达式

	switch time.Now().Weekday() { //  time.Now().Weekday()  获取 今天时星期几
	// 如果时 星期6 或 星期日时 执行
	case time.Saturday, time.Sunday:
		fmt.Println("It 's the weekend")
	// 上面的都不匹配时执行
	default:
		fmt.Println("It 's then weekday")
	}

	t := time.Now() // 获取当前的时间对象
	switch {
	case t.Hour() < 12: //t.hour 获取到当前时间小时数 如果当前时间的小时数小于 12
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	/*
		类型开关：
		声明了一个匿名函数，参数类型为 interface{}
		interface{} 是一种空接口，可以接收任意类型的数据
	*/
	whatAmI := func(i interface{}) {
		switch t := i.(type) { //type 表示将 i 的实际类型进行判断 ,并将类型赋给变量 t
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t) // 使用 %T 打印变量 t 的实际类型
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

}
