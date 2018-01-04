package main

import "fmt"

func printSlice(x []int)  {
	/* len:切片初始长度， cap:切片最大容量，x:打印切片数据 */
	fmt.Printf("len = %d, cap = %d, slice = %v\n",len(x), cap(x), x)
}

func main()  {
	var numbers = make([]int, 3, 5)

	printSlice(numbers)
    /* 空(nil)切片 */
	var slice1 []int

	printSlice(slice1)
	if (slice1 == nil) {
		fmt.Printf("切片是空的")
	}
}