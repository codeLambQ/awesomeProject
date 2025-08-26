package main

import "net/rpc"

func main() {
	// 指定连接的 RPC 服务器
	client, _ := rpc.Dial("tcp", "127.0.0.1:1234")
	var replay string
	//调用哪个方法
	client.Call("HelloService.Hello", "immoc", &replay)
	println(replay)
}
