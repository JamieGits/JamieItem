package main

import "fmt"
/**
 * 数据类型转换
 */
func main()  {
	var a int = 17
	var b int = 5
	var c float32
	c = float32(a)/float32(b)
	fmt.Println(c)
	fmt.Printf("%f\n", c)
}
