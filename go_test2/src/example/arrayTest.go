package main

import "fmt"
/* 定义一个长度为10类型为int的数组arr*/
var arr [10]int
var i, j int
func main()  {
	/* 初始化数组arr */
	for i=0; i<len(arr); i++{
		arr[i] = 100 + i;
	}
	fmt.Printf("%v\n", arr)
	fmt.Println(arr)
    /* 通过遍历输出每个数组元素的值 */
	for j=0; j<len(arr); j++{
		fmt.Printf("arr[%d] = %d\n", j, arr[j])
	}
}
