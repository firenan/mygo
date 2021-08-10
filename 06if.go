package main

import "fmt"

func main() {
	name := "fire"
	/**
	if else 用法
	 */
	if name == "fire" {
		fmt.Println("name is right")
	} else {
		fmt.Println("name is wrong")
	}
	/**
	赋值可以写在if 后
	 */
	if name := "test"; name == "fire" {
		fmt.Println("name is right")
	} else {
		fmt.Println("name is wrong")
	}

	age := 16
	/**
	if else if
	 */
	if (age <= 0) {
		fmt.Println("age illegal")
	} else if (age > 18) {
		fmt.Println("you are a adult!")
	} else {
		fmt.Println("you are a child")
	}
	/**
	Go 里没有三目运算符，所以即使你只需要基本的条件判断，你仍需要使用完整的 if 语句
	 */
}
