package main

func main() {
	//s1 := []int{1, 2, 3}
	//s2 := []int{1, 2, 3, 4, 5}
	// 输出结果 24 24 则代表切片的固定大小为24个字节
	// 切片结构体源码
	//type slice struct {
	//	array unsafe.Pointer  切片指向的数组 8 字节
	//	len   int 切片存在了多少字节数据  8字节
	//	cap   int  切片的容量  8字节
	//}
	//println(unsafe.Sizeof(s1))
	//println(unsafe.Sizeof(s2))

	sa := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// 截取后的切片容量是从数组的第一位开始算的，所以是9
	s3 := sa[1:4]
	println(cap(s3))
	println(len(s3))
	s3 = append(s3, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19)
	println(cap(s3))
	println(len(s3))

	// 切片的扩容
	// 扩容的原理分析
	// 会调用 growslice
	//func growslice(oldPtr unsafe.Pointer, newLen(13), oldCap(9), num int, et *_type) slice {
	//	newcap := nextslicecap(newLen, oldCap)
	//}
	//func nextslicecap(newLen, oldCap int) int {
	//	newcap := oldCap
	// 先对容量进行两倍扩容
	//	doublecap(18) := newcap + newcap
	//  判断扩容后的len是否大于原来两倍的容量，大于的话直接返回 len 作为容量
	//	if newLen(13) > doublecap {
	//		return newLen
	//	}
	//
	//	const threshold = 256
	//	判断旧的容量是否大于256，不大于的话直接返回两倍容量
	//	if oldCap < threshold {
	//		return doublecap
	//	}
	//	for {
	//		以上都不满足的话会扩容到之前的1.25倍
	//		newcap += (newcap + 3*threshold) >> 2
	//		if uint(newcap) >= uint(newLen) {
	//			break
	//		}
	//	}
	//
	//	if newcap <= 0 {
	//		return newLen
	//	}
	//	return newcap
	//}
}
