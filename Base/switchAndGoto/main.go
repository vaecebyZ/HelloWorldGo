package main

import "fmt"

func main() {

	//switch例子

	switch n := 5; n {
	case 1, 2, 3, 4:
		fmt.Println("1")
	case 5:
		fmt.Println("5")
	}

	switch s := 2; {
	case s >= 1:
		fmt.Println("2")
	case s > 2:
		fmt.Println("1")
	case s > 5:
		fmt.Println("5")
	default:
		fmt.Println("233")
	}

	//goto
	//正常情况下不会跳出大循环
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'B' {
				break
			}
			fmt.Printf("%v-%c\n", i, j)
		}
	}

	fmt.Println()

	//使用goto 不止goto可以这样用 break 也能用只不过label必须定义在for switch select代码块上 continue上也可以使用需要定义在for上
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				goto tag //指定跳转
			}
			fmt.Printf("%v-%c\n", i, j)
		}
	}

tag: //lable
	fmt.Print("over")

}
