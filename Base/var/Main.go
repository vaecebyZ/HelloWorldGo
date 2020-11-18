package main

import "fmt"

//正常声明
//var s1 string
//var age int
//var isOk bool

//批量声明(常用)
var (
	name string
	age  int
	isOk bool = true

	//注意：全局声明的变量并不需要一定使用
)

func main() {

	//局部变量声明后必须使用
	name = "Sanka"

	age = 18

	//类型推导加短变量声明（只能在函数内使用）
	gender := "man"

	fmt.Printf("name:%s,age:%d,gender:%s", name, age, gender)
	fmt.Println()
	//打印变量属性
	fmt.Printf("Type:%T", gender)

	//特殊变量：_
	//哑元变量丢弃不想使用的值

}
