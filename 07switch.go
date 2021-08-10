package main

import (
	"fmt"
	"time"
)

func main() {
	/**
	传统方式
	 */
	i := 3
	switch i {
	case 1:
		fmt.Println("the num is 1")
	case 2:
		fmt.Println("the num is 2")
	case 3:
		fmt.Println("the num is 3")

	}
	/**
	在一个 case 语句中，可以使用逗号来分隔多个表达式
	 */
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("today is weekend")
	default:
		fmt.Println("today is weekday")
	}
	/**
	用switch实现if else
	 */
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("morning")
	default:
		println("afternoon")
	}
}
