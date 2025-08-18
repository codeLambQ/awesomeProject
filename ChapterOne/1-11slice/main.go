package main

import "fmt"

func main() {

	// 切片的创建，截取数组
	r := [5]int{1, 2, 3, 4, 5}

	// 截取第1位到第3位
	s := r[1:4]
	fmt.Println(s)

	// 另外的截取方式
	// 截取开头到第3位
	// s1 := r[:4]
	// 截取第1位到结尾
	// s2 := r[ChapterOne:]
	// 截取全部
	// s3 := r[:]

	// 另外一种创建方式
	s4 := []int{1, 2, 3, 4, 5}
	fmt.Println(s4)

	// 切片的访问
	// 直接访问下标
	println(s[1])

	// 访问切片长度
	println("数组长度 = ", len(s))

	// for range 访问
	for i, v := range s {
		fmt.Printf("s[%d] = %d\n", i, v)
	}

	println("数组的追加")
	s = append(s, 7, 8, 9)
	for i, v := range s {
		fmt.Printf("s[%d] = %d\n", i, v)
	}
}
