package main

import (
	"net"
	"net/rpc"
)

type HelloService struct{}

// 创建 HelloService 结构体的方法
func (h *HelloService) Hello(req string, replay *string) error {
	*replay = "hello: " + req
	return nil

}

func main() {
	//注册一个 rpc 服务
	rpc.RegisterName("HelloService", new(HelloService))
	//监听一个端口
	listen, _ := net.Listen("tcp", ":1234")
	// 创建一个 socket 连接
	conn, _ := listen.Accept()
	// 降 socket 绑定到 rpc 上
	rpc.ServeConn(conn)
}
