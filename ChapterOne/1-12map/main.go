package main

import (
	"fmt"
)

func main() {
	// map 的创建
	// map 是 key-value 结构，一个 key 对应一个 value
	// map[keu 的数据类型]value 的数据类型
	// 创建一个空的 map
	m1 := make(map[int]string)
	m1[1] = "123"
	m1[2] = "456"
	m1[3] = "789"

	// 字面量
	m2 := map[int]string{
		1: "123",
		2: "456",
		3: "789",
	}

	// map 遍历
	for k, v := range m2 {
		fmt.Println(k, v)
	}

	// map key 的个数
	println(len(m2))

	// 删除 map key
	delete(m2, 1)
	// 这里删除了 key 为1的元素，但是还是会输出一个空的值
	fmt.Println(m2[1])
	println("1111")

	// 这里创建了两个变量，s 是 key 的值，ok 是用来确定你这个 key 是否存在
	s, ok := m2[1]
	fmt.Println(s)
	fmt.Println(ok)
	if ok {
		println("该 key 存在")
	} else {
		println("该 key 不存在")
	}

	// 并发访问 map，更安全
	//m3 := sync.Map{}
}
