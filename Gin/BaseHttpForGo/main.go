package main

import (
	"fmt"
	"io/ioutil"
	"net/http" //Go自带http库
)

//必须包换http.ResponseWriter 和 http.Request
func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get")
	b, _ := ioutil.ReadFile("./hello.txt")
	_, _ = fmt.Fprintln(w, string(b))
}

func main() {
	//访问地址
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":8989", nil)
	if err != nil {
		fmt.Printf("server error:%v\n", err)
	}
}
