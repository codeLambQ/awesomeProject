package main

import (
	"errors"
	"fmt"
	"strconv"
)

// 自己创建一个函数，来实现 error 的 Error() 方法
func delete1(x, y float64) (float64, error) {
	if y == 0 {
		return 0.0, errors.New("cannot delete")
	}
	return x / y, nil
}

func main() {

	// 可以将 error 的返回值匿名一下，不要了
	i, _ := strconv.Atoi("123")
	fmt.Println(i)

	// 但还如果输入的不是数字的话就会返回一个默认值
	i1, _ := strconv.Atoi("sdasd")
	fmt.Println(i1)

	// 这时候就需要使用 error 的返回值来判断是否有错误
	i2, err := strconv.Atoi("sdasd")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i2)
	}

	f, err := delete1(2, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(f)
	}

	// 输出结果
	//123
	//0
	//strconv.Atoi: parsing "sdasd": invalid syntax
	//cannot delete
}
