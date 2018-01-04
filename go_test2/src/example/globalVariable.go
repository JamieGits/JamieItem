package main

import "fmt"
 /* 声明全局变量 */
 var g int
 var f = 4

func main()  {
	/* 声明局部变量 */
	var a, b int
	var f = 5
	/* 初始化参数 */
	a = 10
	b = 20
	g = a + b
	fmt.Printf("结果: a = %d, b = %d, g = %d\n",a, b, g)
	/* 全局变量和局部变量名称相同时，局部变量优先于全局变量被使用 */
	fmt.Printf("优先选择值为: %d",f)
}



