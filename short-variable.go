package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	// 注意 声明的变量必须使用！！！例如去掉下面的语句，则有语法错误！！！大赞这样的特性！！！
	fmt.Println(i, j, k, c, python, java)
}

/*

短声明变量

在函数中， := 简洁赋值语句在明确类型的地方，可以用于替代 var 定义。

函数外的每个语句都必须以关键字开始（ var 、 func 、等等）， := 结构不能使用在函数外。

 */