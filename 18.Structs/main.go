package main

import "fmt"

// 声明一个结构体
type person struct {
	name string
	age  int
}

// 用于初始化结构体对象的方法,放回了一个结构体的地址
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	// 初始化结构体
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Fred"})
	fmt.Println(person{
		name: "Ann", age: 40})
	fmt.Println(newPerson("Jon"))

	// 创建一个 结构体对象
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
	sp := &s // 将s 结构体的地址赋值给sp实例
	// 通过指针的方式去访问和修改person 实例的字段
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
	fmt.Println(s.age)

	// 匿名结构体
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}
