package main

type myJob struct {
	name   string
	adress string
	year   int
}

// 方法是和结构体绑定的，传入的参数的时候也是值传递
//func (m1 myJob) updateName(newName string) {
//	m1.name = newName
//}

// 可以传入结构体的指针
func (m1 *myJob) updateName(newName string) {
	m1.name = newName
}

func main() {
	//m1 := myJob{"ops", "beijing", 2}
	//// 输出结果：ops 可以验证为值传递，不会修改原来的值
	//m1.updateName("devops")
	//println(m1.name)

	m1 := myJob{"ops", "beijing", 2}
	// 输出结果：devops 修改了原来的值
	m1.updateName("devops")
	println(m1.name)
}
