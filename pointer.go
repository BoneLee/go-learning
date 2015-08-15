package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

/*
// 喜欢这个特性 扬长避短 因为这样就获得了C++的高效和灵活性
// ！！！函数是按照值传递！！！这样更不容易出错！！！

指针

Go 具有指针。 指针保存了变量的内存地址。

类型 *T 是指向类型 T 的值的指针。其零值是 nil 。

var p *int

& 符号会生成一个指向其作用对象的指针。

i := 42
p = &i

* 符号表示指针指向的底层的值。

fmt.Println(*p) // 通过指针 p 读取 i
*p = 21         // 通过指针 p 设置 i

这也就是通常所说的“间接引用”或“非直接引用”。

与 C 不同，Go 没有指针运算。

 */