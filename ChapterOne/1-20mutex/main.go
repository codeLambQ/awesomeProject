package main

import (
	"sync"
	"time"
)

// 解决 1000 goroutine 对变量操作的并发问题

// 创建一个互斥锁
var m = sync.Mutex{}

func add(p *int) {
	// 操作之前先加锁
	m.Lock()
	// 注册解锁操作，这样就会在这个函数结束之前执行解锁操作，会让程序更安全
	defer m.Unlock()
	*p++
}

func main() {
	i := 0
	for j := 0; j < 1000; j++ {
		go add(&i)
	}
	//group := sync.WaitGroup{}
	//group.Wait()
	time.Sleep(6 * time.Second)
	println(i)
}
