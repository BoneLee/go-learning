package main

import "fmt"

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java) //0 false false false
	// 可以看到，它是有默认值的，int的默认值是0，这和我倡导的c++语言规范是多么吻合啊
}

/*

变量

var 语句定义了一个变量的列表；跟函数的参数列表一样，类型在后面。

就像在这个例子中看到的一样， var 语句可以定义在包或函数级别。

 */