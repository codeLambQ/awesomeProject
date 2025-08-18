package main

func main() {

	// 创建一个变量为 string 类型
	// var 是用来声明这个 a 是一个新的变量
	var a string = "hello"
	println("a = ", a)

	// 也可以同时创建多个相同类型的变量，需要使用英文逗号 , 分割
	// 这里没有写数据的类型，因为 go 会根据变量初始化的值来进行推断
	var b, c = 1, 2
	println("b = ", b, "c = ", c)

	// 也可以不用写 var,创建变量，使用 := 来代替
	e := "e"
	println("e = ", e)

	// 在运行过程中不可以更改变量的数据类型
	//e = ChapterOne  这样是错误的

	//f = "ll"  不可以这样创建新的变量，这个语法是给已经创建的变量赋值

	// 初始化变量的时候可以不给变量赋值，但是必须给变量确定数据类型
	//var g 这样是错的，go 没有办法来推断你是哪种数据类型
	// 这样创建变量 go 会给变量一个默认值
	var g float32
	println("g = ", g)

}
