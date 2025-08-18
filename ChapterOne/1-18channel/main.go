package main

import (
	"fmt"
	"time"
)

// 异步处理字符串的函数
func processString(c chan string) {
	for s := range c {
		println("processString start process", s)
		time.Sleep(10 * time.Second)
		println("processString end process", s)
	}
}

func main() {
	// channel 的简单使用
	// 声明无缓存的 channel
	//c := make(chan string)
	// 声明有缓存的 channel
	//c := make(chan string,2)
	// 往 channel 里面塞数据 c <- "sdasd"
	// 从 channel 里面取出数据 s := <- c
	// 错误用法, 创建了一个无缓存的 channel，并且放了同一个 Goroutine 中使用
	//c := make(chan string)
	//c <- "moody"
	//<- c

	c := make(chan string, 5)
	go processString(c)
	var s string
	println("main() start")
	for true {
		fmt.Scanln(&s)
		c <- s
	}

	// 输出结果
	//main() start
	//a
	//processString start process a
	//b
	//c
	//d
	//e
	//processString end process a
	//processString start process b
	//processString end process b
	//processString start process c
	//processString end process c
	//processString start process d
	//processString end process d
	//processString start process e
	//processString end process e

}
