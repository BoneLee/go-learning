package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)

	/* EQUAL
	var a int
	a = 100
	*/

	var a = 100;

	/* NOT allowed
	var a
	a = 100
	*/

	// NOT allowed
	// k := 3 /*k declared and not used */

	fmt.Println(a)
}

/*

初始化变量

变量定义可以包含初始值，每个变量对应一个。

如果初始化是使用表达式，则可以省略类型；变量从初始值中获得类型。

 */