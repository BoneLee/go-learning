package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the playground!")

	fmt.Println("The time is", time.Now())
}

/*
    在 Playground 中，时间从 2009-11-10 23:00:00 UTC（了解这个日期的重要含义是留给读者的练习）。给程序赋确定性的输出使得缓存程序的输出变得容易。
    对于运行时间、CPU 和内存的使用同样也有限制，并且程序不能访问外部网络中的主机。
 */