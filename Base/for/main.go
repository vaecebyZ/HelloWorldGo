package main

import "fmt"

func main() {
	//for-------------------------------------

	//	基本格式

	for i := 0; i < 10; i++ {
		fmt.Println(i) //0-9
	}
	//变种1 作用域外声明i
	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}
	//变种2  省略 执行
	for i < 20 {
		fmt.Println(i)
		i++
	}

	//变种3 无限循环
	//for {
	//
	//}

	//变种4 for range 返回键值对 常用

	s := "你好SANKA"

	for i, v := range s {
		//i索引 v值
		fmt.Printf("%d,%c", i, v) //
	}
	//例子
	//当i==5时候跳出
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Printf("出来了%v", i)
			break
		}
	}

	//但i==5时候跳过
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Printf("%v\n", i)
	}

}
