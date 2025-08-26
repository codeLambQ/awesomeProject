package main

import (
	"fmt"
	"net/http"
)

// 最简单的HTTP实现

// 创建一个函数，不论请求是啥都返回 hello world

func hello(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "Hello World")
}
func main() {
	// 注册一个路径
	http.HandleFunc("/hello", hello)
	// 注册一个端口
	http.ListenAndServe(":8080", nil)
}
