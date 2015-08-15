package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

/*
// 我的理解它就是类似在内存stack里压入里一个函数指针，当外层函数执行完成，代码继续用弹出函数执行！

defer 栈

延迟的函数调用被压入一个栈中。当函数返回时， 会按照后进先出的顺序调用被延迟的函数调用。

阅读博文了解更多关于 defer 语句的信息。

 */