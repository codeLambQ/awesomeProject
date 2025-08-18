package main

import (
	"fmt"
	"reflect"
)

type myLesson struct {
	id     int
	name   string
	rating float64
}

// 给结构体创建一个方法
func (ml myLesson) printName() {
	fmt.Println(ml.name)
}

// 反射: 通过数据可以查看元数据
func main() {
	m := myLesson{555, "MySql", 10.0}
	typeOf := reflect.TypeOf(m)
	fieldByName, has := typeOf.FieldByName("rating")
	println(fieldByName.PkgPath)
	println(has)

	methodByName, hash := typeOf.MethodByName("printName")
	println(methodByName.Name)
	println(hash)
}
