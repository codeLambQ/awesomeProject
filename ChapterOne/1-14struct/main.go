package main

import "fmt"

func main() {
	// 结构体
	type myJob struct {
		name   string
		adress string
		year   int
	}

	// 创建一个结构体变量，并初始化
	m1 := myJob{name: "运维", adress: "北京", year: 3}
	// 另一种方法,按照顺序来的话不需要指定路径
	// m1 := myJob{"运维", "北京", 3}

	// 访问
	fmt.Println(m1)
	println(m1.year)

	// 结构体组合，相当于进程
	type myJob2 struct {
		myJob
		id int
	}

	m2 := myJob2{}
	println(m2.id, m2.myJob.name, m2.adress)
}
