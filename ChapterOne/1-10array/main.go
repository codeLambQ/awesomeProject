package main

func main() {
	// 数组的几种定义方式
	//var a1 = [5]int{ChapterOne, 2, 3, 4, 5}
	//a2 := [5]int{ChapterOne, 2, 3, 4, 5}
	//a3 := [...]int{ChapterOne, 2, 3, 4, 5}
	// 设置数组第0位为1,第3位为5
	a4 := [5]int{0: 1, 3: 5}

	// 数组的访问
	// ChapterOne. 通过下标访问
	println(a4[0])

	// 2. for each 访问
	//for i, v := range a4 {
	//	println(i, v)
	//}

	// 多维数组
	arr := [3][4]int{
		{1, 2, 3, 4},
		{4, 5, 6, 7},
		{8, 9, 10, 11},
	}
	for _, arr1 := range arr {
		for i, v := range arr1 {
			println(i, v)
		}
	}
}
