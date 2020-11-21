package main

import "fmt"

//if条件判断
func main(){

	//if--------------------------------------------------------
	age := 18
	if age>=18 {
		fmt.Println("成年了")
	}else {
		fmt.Println("去写作业吧")
	}

	fmt.Println("")

	if age>35{
		fmt.Println("中年")
	}else if age>=18{
		fmt.Println("青年")
	}else {
		fmt.Println("学习去吧")
	}

	fmt.Println("")

	//特殊写法
	if aage:=18 ;aage >= 18{
		fmt.Println("成年了")
	}

}
