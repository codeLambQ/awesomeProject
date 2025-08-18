package main

import (
	"fmt"
	"time"
)

func remindMeLater(t time.Duration, notice string) {
	time.Sleep(t)
	fmt.Println(notice)
}

func main() {
	// 执行这个之后会发现执行完吃饭提醒之后，间隔5秒后才会执行睡觉
	//remindMeLater(3*time.Second, "吃饭")
	//remindMeLater(5*time.Second, "睡觉")

	// 如果想要并发的去运行这两行函数，只需要在函数前面加一个 go 即可
	// 这样 go 语言就会单独的将这两行函数放在不通的 Goroutine 中去运行，go 语言所有的函数都是放在 Goroutine 去运行的
	// 执行完后会发现先执行了 println("main() end")，可以论证 main 函数也是放在了 Goroutine 中运行的
	//main() end
	//吃饭
	//睡觉
	//8
	go remindMeLater(3*time.Second, "吃饭")
	go remindMeLater(5*time.Second, "睡觉")

	println("main() end")
	fmt.Scanf("%d")
}

// 总结 什么是 Goroutine
//ChapterOne. 并发的执行函数，如果不是函数可以包装成匿名函数
//go func() {
//	println("1111")
//}()
// 2.Go 中的所有逻辑都是放在 Goroutine 中运行的
