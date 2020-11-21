package main

import (
	"fmt"
	"strings"
)

//字符串操作utf-8实现
//“”为字符串
//‘’是字符

// \转义符更
// \r回车符号
// \n换行
// \t制表

func main() {

	s1 := "字符串" //字符串

	s2 := '号' //字符

	fmt.Printf("%v%v\n", s1, s2)

	s3 := `		sanka
				vaecebyz,
				不以物喜，不以己悲
		  ` //保持格式

	fmt.Printf(s3 + "\n")

	//---------------------------------------查看长度

	fmt.Print(len(s1)) //9 一个汉字3个字节
	fmt.Print("\n")
	fmt.Println(len("S")) //1 一个字母占用一个字节

	//---------------------------------------字符拼接

	name := "sanka"
	bas := "vaecebyz"

	fmt.Println(name + " " + bas)

	sambas := fmt.Sprintf("40%s%s", name, bas) //可以把字符返回 不接收的话是不会打印的

	fmt.Sprintf("41%s%s", name, bas) //不会打印

	fmt.Println(sambas)

	fmt.Printf("%v%v", name, bas) //直接打印

	fmt.Print("\n")

	//------------------------------------字符串分割

	s4 := `https://vaecebyz.github.io` //`不需要转译

	fmt.Printf("%T:%v\n", strings.Split(s4, "/"), strings.Split(s4, "/")) //按照/线分割成数组

	//------------------------------------是否包含

	fmt.Println(strings.Contains(s4, "https")) //true

	fmt.Println(strings.Contains(s4, "2")) //false

	//------------------------------------开头和结尾

	fmt.Println(strings.HasPrefix(s4, "https")) //true 以https开头

	fmt.Println(strings.HasSuffix(s4, ".io")) //true 以.io结尾

	//------------------------------------获取指定字符串位置

	fmt.Println(strings.Index(s4, "i")) //18  开始位置

	fmt.Println(strings.LastIndex(s4, "i")) //24 结束的位置

	//------------------------------------拼接

	fmt.Println(strings.Join(strings.Split(s4, "/"), "+")) //以/拆分以+号拼接

	//------------------------------------len遍历

	//len()求的是byte字节数量

	//非ascii码类型字符为rune类型

	//ascii码类型的为 byte 或者叫 unit8

	s5 := "Sanka さんか"

	for _, c := range s5 {
		fmt.Printf("%c\n", c)
	}

	//-----------------------------------字符串修改

	s6 := "你好吗？"

	//s6[0] = "1" 无法这样修改

	s66 := []rune(s6) //把s6转换成rune切片 -> ['你' ’好‘ ’吗‘ ’？‘]

	s66[0] = '我' //只能用字符类型

	fmt.Printf("%T,%T\n", "你", '你') //rune 是int32的别名 byte 是 uint8的别名

	fmt.Println(string(s66)) //转换成字符串

}
