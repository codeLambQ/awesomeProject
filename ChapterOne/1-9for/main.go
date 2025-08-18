package main

import "fmt"

func main() {
	// for 语言语法
	// for 初始化; 条件; 对初始化变量进行处理 {}

	//for i := 0; i < 5; i++ {
	//	print(i)
	//}

	// 还可以将 i := 0; i < 5; i++ 进行分开
	//j := 0
	//for j < 5 {
	//	print(j)
	//	j++
	//}

	// break 和 continue 用法
	// break 直接退出循环
	// continue 退出本次循环，继续执行下面的玄幻

	//for i := 0; i < 5; i++ {
	//	if i == 3 {
	//		break
	//	}
	//	print(i)
	//	// 输出结果012
	//}

	//for i := 0; i < 5; i++ {
	//	if i == 3 {
	//		continue
	//	}
	//	print(i)
	//	// 输出结果0124
	//}

	// for each
	arr := []int{1, 2, 3, 4, 5}
	for i, v := range arr { // 这里 i 代表数组的索引，v 代表数组的值
		fmt.Printf("%d = %d \n", i, v)
	}
}
