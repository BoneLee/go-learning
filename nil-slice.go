package main

import "fmt"

func main() {
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}
}
// 由此看数组的初始化是nil还算正常，因为c++里数组是指针，指针初始化为空正常！
/*

nil slice

slice 的零值是 nil 。

一个 nil 的 slice 的长度和容量是 0。

 */