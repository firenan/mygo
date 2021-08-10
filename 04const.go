package main

import (
	"fmt"
	"math"
)

/**
const 语句可以出现在任何 var 语句可以出现的地方
 */
const s = "constant"
func main()  {
	fmt.Println(s)

	const n = 1000
	/**
	常数表达式可以执行任意精度的运算
	 */
	const d = 3e20 / n
	fmt.Println(int64(d))
	/**
	数值型常量是没有确定的类型的，直到它们被给定了一个类型，比如说一次显示的类型转化。
	这里的 math.Sin函数需要一个 float64 的参数。
	 */
	fmt.Println(math.Sin(n))
}
