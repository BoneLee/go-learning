package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}

/*

函数

函数可以没有参数或接受多个参数。

在这个例子中， add 接受两个 int 类型的参数。

注意类型在变量名 之后 。

（参考 这篇关于 Go 语法定义的文章了解类型以这种形式出现的原因。）

因为go的哲学是可读性更高的语法，将变量类型防止前面的做法给人一种喧宾夺主的感觉，因为重要的东西明显应该
放在前面。函数名字肯定比返回值更重要。
 */