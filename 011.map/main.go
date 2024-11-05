package main

import (
	"fmt"
)

func main() {
	//创建一个空make

	m := make(map[string]int)

	// 使用典型的 name[key] = val 语法设置键/值对
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map", m)

	// 通过key 获取值
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	v3 := m["k2"]
	fmt.Println("v3:", v3)

	// 通过 len 获取到 map 中的键值对数量
	fmt.Println("len:", len(m))

	// 通过delete 可以 在map 删除键值对
	delete(m, "k2")
	fmt.Println("map:", m)

	// 通过clear 可以删除map 中所有键值对
	clear(m)
	fmt.Println("map:", m)

	//从 map 获取值时，可选的第二个返回值指示 map 中是否存在键,当不需要值的本事时可以通过_忽然它
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	//在同一行中声明和初始化一个新map
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	// 也可以通过 Equal去 判断两个 map 是否相同
	n2 := map[string]int{"foo": 1, "bar": 2}
	if test(n, n2) {
		fmt.Println("n == n2")
	}
}

func test[M1, M2 ~map[k]v, k, v comparable](m1 M1, m2 M2) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v1 := range m1 {
		if v2, ok := m2[k]; !ok || v1 != v2 {
			return false
		}
	}
	return true
}
