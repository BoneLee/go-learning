package main

import "fmt"

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	// 我有疑问，常量应该也有数值类型的啊 推导其类型的时候如何定？
//	fmt.Println(Big)  //错误：constant 1267650600228229401496703205376 overflows int
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

/*

数值常量

数值常量是高精度的 值 。

一个未指定类型的常量由上下文来决定其类型。

也尝试一下输出 needInt(Big) 吧。

 */