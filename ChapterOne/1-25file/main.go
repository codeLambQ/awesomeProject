package main

import (
	"os"
)

func readFile(file *os.File) {
	// 读取文件需要先打开文件
	fopen, _ := os.Open(file.Name())
	bytes := make([]byte, 1024)

	// 循环读取
	for {
		// 每次只能读取一行内容
		n, _ := fopen.Read(bytes)
		if n == 0 {
			break
		}
	}
	defer func(fopen *os.File) {
		err := fopen.Close()
		if err != nil {
			err.Error()
		}
	}(fopen)

	// 读取内容
	_, werr := os.Stdout.Write(bytes)
	if werr != nil {
		werr.Error()
	}

}

func writeFile(f *os.File, s string) {
	_, err := f.WriteString(s)

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			err.Error()
		}
	}(f)

	if err != nil {
		err.Error()
	}

}

func createName(name string) *os.File {
	file, err := os.Create(name)

	if err != nil {
		err.Error()
	} else {
		return file
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			err.Error()
		}
	}(file)
	return file
}

func main() {
	// 创建文件夹
	_ = os.Mkdir("ql", 0777)

	// 创建多层级目录
	_ = os.MkdirAll("ql\\test1\\test2", 0700)

	// 删除目录，以及目录里面的内容
	_ = os.RemoveAll("ql")

	// 创建文件，go 语言创建文件后会打开文件，必须要关闭，否则其余人不可以用改文件
	file := createName("ChapterOne-25file\\ql.txt")

	// 写入内容
	writeFile(file, "ql\r\nlll")

	// 读取文件内容
	readFile(file)

	// 删除文件
	//_ = os.RemoveAll("ql.txt")
}
