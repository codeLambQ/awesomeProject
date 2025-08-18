package main

func main() {
	// 指针
	var i int = 9 // 定义的变量是存储在内存里面的
	// 指针用来定位到变量在内存中的位置 &i 来获取
	// *int 代表整型的指针类型
	var p *int = &i
	// 整型的指针的指针类型
	var pp **int = &p

	// 输出看一下
	println(i)
	println(&i)
	println(p)
	println(&p)
	println(pp)

	// 用来获取该内存地址的原值
	println(*p)

	// Go 语言的方法参数都是值传递，方法里面对参数的操作不会影响原值
	println(i)
	add(i)
	println(i)

	// 将在内存中的地址给到方法的话就可以修改原值
	println(i)
	padd(p)
	println(i)
}

func add(j int) {
	j++
}

func padd(p *int) {
	//j++
	*p++
}
