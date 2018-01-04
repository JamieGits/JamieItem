package main

import "fmt"

func main()  {
	/* 声明局部变量 */
	var a, b, c int
    /* 初始化参数 */
	a = 10
	b = 20
	c = a + b

   var	flag bool

	fmt.Printf("a = %b\n",a)  //转换为十进制数值
	fmt.Printf("a = %o\n",a)  //转换为八进制数值
	fmt.Printf("a = %x\n",a)  //转换为小写的十六进制
	fmt.Printf("a = %X\n",a)  //转换为大写的十六进制
	fmt.Printf("flag = %t\n",flag) //输出布尔值
	fmt.Printf("输出值flag的类型是: %T\n",flag) //输出值的类型
	fmt.Printf("结果: a = %d, b = %d, c = %d\n", a, b, c)
}
