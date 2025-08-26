package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type LessonServiceImpl struct {
	UnimplementedLessonServiceServer
}

func (l *LessonServiceImpl) GetLesson(context.Context, *LessonId) (*Lesson, error) {
	lesson := Lesson{}
	lesson.Id = 123
	lesson.Name = "mysql"
	lesson.Rating = 10.0
	return &lesson, nil

}
func main() {
	lessonServiceClient := LessonServiceImpl{}
	// 注册一个端口
	listen, _ := net.Listen("tcp", ":8080")
	// 创建一个 grpc 服务
	grpcServer := grpc.NewServer()
	// 注册服务端的 grpc 服务
	RegisterLessonServiceServer(grpcServer, &lessonServiceClient)
	// Serve 会通过监听器 lis 接收传入的连接，为每个连接创建一个新的 ServerTransport 和服务 goroutine
	err := grpcServer.Serve(listen)

	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}
