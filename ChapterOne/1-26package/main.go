package main

import (
	"awesomeProject/1-26package/imooc"
	"fmt"
)

func main() {
	// 在同一个 package 下的函数、结构体、方法可以直接引用，不需要导入，并且名字可以随意
	fmt.Println(myLesson{1, "222"})

	// 不同 package 下的函数、结构体、方法不可以直接引用，需要使用 import 关键字先引入
	println(imooc.MyLesson{2, "222"})
}
