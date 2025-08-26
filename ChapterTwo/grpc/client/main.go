package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"time"
)

func main() {
	// 建立 rpc 通道
	conn, _ := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure(), grpc.WithBlock())
	// 注册关闭连接
	defer conn.Close()
	//创建一个链接
	client := NewLessonServiceClient(conn)

	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second)
	defer cancelFunc()
	lessonId := LessonId{Id: 515}
	// 调用方法
	lesson, _ := client.GetLesson(timeout, &lessonId)

	fmt.Println(lesson)
}
