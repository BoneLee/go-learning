package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

/*
不知道内在实现原理是啥？

defer
vt.
使推迟； 使延期； 拖延，推迟； [军]使延期入伍

vi.
服从； 延期； 推迟


defer 语句会延迟函数的执行直到上层函数返回。

延迟调用的参数会立刻生成，但是在上层函数返回前函数都不会被调用。

 */