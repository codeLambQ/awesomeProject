package main

import (
	context "context"
	"github.com/go-kratos/kratos/v2/transport/http"
	"net"
)

type LessonServiceHttpImpl struct{}

func (l *LessonServiceHttpImpl) GetLesson(ctx context.Context, id *LessonId) (*Lesson, error) {
	lesson := &Lesson{}
	lesson.Id = 535
	lesson.Name = "Redis"
	lesson.Rating = 10.0
	return lesson, nil
}

func main() {
	// 生成 socket
	listen, _ := net.Listen("tcp", "0.0.0.0:9090")
	// 创建一个 http server
	httpServer := http.NewServer()
	// 注册一个 http server
	RegisterLessonServiceHTTPServer(httpServer, &LessonServiceHttpImpl{})
	// Serve 会通过监听器 l 接收传入的连接，并为每个连接创建一个新的服务 goroutine
	httpServer.Serve(listen)
}
