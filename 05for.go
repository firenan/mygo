package main

import "fmt"

func main() {
	i := 1
	/**
	单个循环条件
	 */
	for i <= 3 {
		fmt.Println(i)
		i++
	}
	/**
	精点循环写法
	 */
	for j := 1; j <= 4; j++ {
		fmt.Println(j)
	}
	/**
	死循环
	 */
	for {
		fmt.Println(1)
	}
}
