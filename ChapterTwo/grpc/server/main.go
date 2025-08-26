package main

import (
	"context"
	"google.golang.org/grpc"
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
	listen, _ := net.Listen("tcp", ":8080")
	grpcServer := grpc.NewServer()
	RegisterLessonServiceServer(grpcServer, &lessonServiceClient)

	grpcServer.Serve(listen)
}
