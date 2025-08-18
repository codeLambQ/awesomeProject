package main

func main() {
	println(add(1, 2))

	println(swith_ql(1, 2))

	println(arr_int(1, 2, 3))

	var t = add
	println("函数作为变量值：", t(2, 3))

	println("函数作为函数的返回值: ", function())

	// 值传递测试
	var j = 1
	plus(j)
	println("j 的值并没有改变: ", j)
}

// 函数的定义
// func 函数名(函数的入参) 函数返回体 {
//	函数体
//}

// ChapterOne. 基础函数
func add(i, j int) int {
	return i + j
}

// 2. 函数可以有多个返回值
func swith_ql(i, j int) (int, int) {
	return j, i
}

// 3. 函数的可变参数，相当于一个数组
func arr_int(i ...int) int {
	return i[0]
}

// 4. 函数作为变量
// var t = add  这样 t 就变成了函数，和 add 函数的用法一直，相当于是引用

// 5. 函数作为函数的返回值或者参数
func function() func(int, int) int {
	return func(i, j int) int {
		return 1 + 1
	}
}

// 6. 值传递，如果将变量作为参数传给函数，不会影响变量本来的值，只是将变量当时的值传给函数
func plus(i int) {
	i++
}
