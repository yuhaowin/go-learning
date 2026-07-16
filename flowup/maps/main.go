package main

import (
	"fmt"
)

func main() {

	m := map[string]string{
		"name": "yuhao",
	}

	m2 := make(map[string]int, 100) // empty map

	var m3 map[string]int // nil Zero Value 可以安全的使用

	fmt.Println(m, m2, m3)

	fmt.Println(len(m), len(m2))

	for k, v := range m {
		fmt.Println(k, v)
	}

	// key 不存在时，获得 Value 类型的初始值 Zero Value
	s, ok := m["name"]

	fmt.Println(s, ok)

	if s, ok := m["name"]; ok {
		fmt.Println(s)
	}

	delete(m, "name")

}
