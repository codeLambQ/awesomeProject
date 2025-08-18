package main

// 泛型的定义
func min1[T number](a, b T) T {
	if a >= b {
		return b
	}
	return a
}

// 接口也可以用来创建泛型
type number interface {
	~int | ~float64
}

func main() {
	println(min1(10, 5.0))
}
