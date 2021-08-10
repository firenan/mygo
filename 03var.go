package main

import "fmt"

func main() {
	/**
	var 声明 1 个或者多个变量。
	 */
	var a = "string"
	fmt.Println(a)

	var b, c = 1, 2
	fmt.Println(b, c)

	var d, e= true, false

	fmt.Println(d, e)

	/**
	声明变量且没有给出对应的初始值时，变量将会初始化为零值 。例如，一个 int 的零值是 0。
	 */
	var f int
	var i bool
	fmt.Println(f, i)
	/**
	:= 语句是申明并初始化变量的简写，g := "short" 等同于 var g string = "short"。
	 */
	g := "short"
	h := 123
	fmt.Println(g, h)

}
