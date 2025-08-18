package main

// 定义一个接口
// Go 语言中的接口都是隐式接口，可以在不修改源代码的情况下抽象出接口
type wuFanBanLv interface {
	play()
}

// 定义结构体实现接口方法
type myLession struct {
	name string
}

func (m1 myLession) play() {
	println("playing ", m1.name)
}

func main() {
	m1 := myLession{"云原生课程"}
	chifan(m1)
}

func chifan(w wuFanBanLv) {
	w.play()
}
