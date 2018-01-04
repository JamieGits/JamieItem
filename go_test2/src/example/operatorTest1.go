package main

import "fmt"

func main()  {

	/*其他运算符*/
	var a  int = 4
	var b  int32
	var c  float32
	var ptr  *int

	/* 运算符实例*/
	fmt.Printf("第一行 -a变量类型为 = %T\n",a)
	fmt.Printf("第二行 -b变量类型为 = %T\n",b)
	fmt.Printf("第三行 -c变量类型 = %T\n",c)

	/* &和 * 运算符实例 */
	ptr = &a  //ptr包含了a变量的地址
	fmt.Printf("a的值为 %d\n",a)
	fmt.Printf("ptr为 %d\n",*ptr)
}
