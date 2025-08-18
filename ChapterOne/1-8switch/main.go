package main

func main() {
	// switch 语法
	var grade rune = 'C'

	// 默认情况下 switch 只会执行匹配的那个逻辑，不会执行下面的 case 的内容
	// 如果要执行的话可以使用 fallthrough
	// 输出结果
	//C
	//B
	switch grade {
	case 'C':
		println("C")
		fallthrough
	case 'B':
		println("B")
	case 'A':
		println("A")
	default:
		println("你这不对啊")

	}
}
