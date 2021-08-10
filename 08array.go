package main

import "fmt"

func main() {
	var a [5] int

	fmt.Println("arr int", a)
	a[4] = 100
	fmt.Println("arr1", a)
	fmt.Println("len a", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("arr b", b)

	c := b[1]
	fmt.Println(c)

	var d [2][3] int

	fmt.Println("二维数组", d)
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			d[i][j] = i + j
		}
	}

	fmt.Println("二维数组赋值", d)

	d[0][0] = 2

	fmt.Println("修改二维数组中的值", d)
}
