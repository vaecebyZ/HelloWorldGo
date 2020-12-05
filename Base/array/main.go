package main

import "fmt"

//数组
func main() {
	//必须指定存放的类型

	//定义一个长度为3的bool数组

	var ai [3]bool  //[true , false,true]
	var aii [4]bool //[true,false , true]

	//数组的长度是 类型的一部分所以这两个数组不能进行比较
	fmt.Printf("ai:%T,aii:%T\n", ai, aii)

	//数组的初始化
	//如果不初始化所有内容默认为0 0在bool里为false

	//方法一
	ai = [3]bool{true, true, true}
	fmt.Printf("%v \n", ai)

	//方法二
	//aiii := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//根据初始值自动推断长度
	aiii := [...]int{6, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(aiii)

	//方法三
	//根据索初始化
	aiv := [3]int{0: 1, 2: 2}
	fmt.Println(aiv)

	//数组遍历
	//根据索引遍历
	for i := 0; i < len(aiv); i++ {
		fmt.Println(aiv[i])
	}

	//for range遍历

	for _, v := range aiii {
		fmt.Println(v)
	}

	//多维数组
	//[[1,2],[3,4]]
	av := [2][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
	}
	fmt.Println(av)

	//多维遍历 双重for _为下标 并不使用所以丢弃
	for _, v := range av {
		fmt.Println(v)
		for _, vs := range v {
			fmt.Println(vs)
		}
	}

	//数组是值类型
	//ex
	avi := [...]int{1, 2, 3}
	avii := avi //直接复制值
	avii[0] = 12
	fmt.Printf("avi:%v,avii:%v\n", avi, avii) //结果不会影响之前的数组
}
