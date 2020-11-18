package main

import "fmt"

const (
	pi   = 3
	e    = 2
	code = 400
	code2
	code3
	code4
	//code234为上一行的常量的声明值
)

//iota常量计数器
//当他出现的时候为0
//随后每出现一次常量声明iota++
const (
	n1 = iota //0
	n2        //1
	n3        //2
	n4 = 1    //1
	n5 = iota //4
	_         //丢弃 iota = 5
	n6        //6
)

func main() {
	// pi = 4 无法通过编译 定义后就无法修改
	//全部为400
	fmt.Println(code, code2, code3, code4)
	//0 1 2 1 4 6
	fmt.Println(n1, n2, n3, n4, n5, n6)
}
