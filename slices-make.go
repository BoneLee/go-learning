package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

// 发现没有 其实go里也是借鉴了c++的特性
// 数组是有容量这个东西的
// 越来越发现go这个东西就是C++的变种了，只是google的家伙加速了编译过程，引入了一些更高级的特性而已

/*

构造 slice

slice 由函数 make 创建。这会分配一个零长度的数组并且返回一个 slice 指向这个数组：

a := make([]int, 5)  // len(a)=5

为了指定容量，可传递第三个参数到 make：

b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4

 */