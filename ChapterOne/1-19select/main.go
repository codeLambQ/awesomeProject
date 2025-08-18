package main

import "time"

// select 用来取出准备好的 channel 的数据

// 公司年会，谁先回答输出谁的答案
func sender1(c chan string) {
	time.Sleep(7 * time.Second)
	c <- "i am sender ChapterOne"
}

func sender2(c chan string) {
	time.Sleep(5 * time.Second)
	c <- "i am sender 2"
}

func sender3(c chan string) {
	time.Sleep(3 * time.Second)
	c <- "i am sender 3"
}

func main() {
	// 创建三个 channel
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	go sender1(c1)
	go sender2(c2)
	go sender3(c3)

	select {
	case s1 := <-c1:
		println(s1)
	case s2 := <-c2:
		println(s2)
	case s3 := <-c3:
		println(s3)
	}

}
