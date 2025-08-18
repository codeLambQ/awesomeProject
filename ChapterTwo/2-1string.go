package main

import (
	"unicode/utf8"
	"unsafe"
)

func main() {

	// 本质上是结构体的一个值
	s := "慕课网"

	// 字符串在内存中初始大小
	// 输出结果 16
	// 字符串的结构体
	//type stringStruct struct {
	//  定义一个指针变量，指向内存中的地址
	//	str unsafe.Pointer  8字节
	//	字符串的长度   8字节
	//	len int
	//}
	println(unsafe.Sizeof(s))
	println(unsafe.Sizeof(int64(2)))

	// 字符串数组的长度
	// 输出结果 9
	println(len(s))

	// 字符串的字符数
	// 输出结果 3
	println(utf8.RuneCountInString(s))
}
