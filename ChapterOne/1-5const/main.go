package main

func main() {
	// 常量：使用关键字 const 创建，初始化的时候必须赋值，后续要不可以改变值
	// 可以显示指定数据类型
	const a int = 100
	// 也可以不指定让 go 来进行推断
	const b = 200

	// 常量块，可以更方便的创建一组常量
	const (
		a1 = "ll"
		b1 = 23
		c1 = 24
	)

	// iota 可以自增的常量
	const (
		//a2 = iota
		//b2 = iota
		//c2 = iota
		// 另外一种写法,必须是连续的
		a2 = iota
		b2
		c2
	)
	println(a2, b2, c2)
	// 输出结果 012, 从输出结果上可以很直观的看出来 iota 进行自增了

	const (
		a3 = iota
		b3 = "ll"
		c3 = iota
		d3
	)
	println(a3, b3, c3, d3)
	// 输出结果：0 ll 2 3，从输出结果上可以看出来，b3 = ll 这里虽然没有使用 iota，但是还是进行了自增
}
