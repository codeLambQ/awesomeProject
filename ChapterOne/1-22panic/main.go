package main

// 自己创建一个函数，来实现 error 的 Error() 方法
// go 语言中的 error 不够强势，这时候需要引入 panic 机制来解决这种问题
// 改造一下代码
func delete1(x, y int) int {
	if y == 0 {
		//return 0.0, errors.New("cannot delete")
		panic("cannot divide by zero")
	}
	return x / y
}

func divideByZero(x, y int) int {
	// 这种写法的话不会 panic 进行处理，用户可能不知道为什么出现这个问题，来修改一下这种方法
	//defer func() {
	//	recover()
	//}()

	defer func() {
		r := recover()
		if r != nil {
			// 表示出现了问题，来处理一下
			println(r.(string))
		}
	}()
	result := delete1(x, y)
	return result
}

func main() {
	// 如果是这种写法的话遇到 panic 的时候会结束掉函数的运行，不会执行后面的逻辑
	// 解决办法，创建出来一个方法，方法里面来处理 panic
	result := divideByZero(10, 0)
	println(result)

}
