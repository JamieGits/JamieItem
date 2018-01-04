package main

import "fmt"

type MyPoint struct{
	x int
	y int
}
//如果方法的参数是值，那么按照传值得方式，方法内部对struct改动无法作用在外部变量上
func printFunValue(p MyPoint)  {
	p.x = 1
	p.y = 2
	fmt.Printf("-> %v\n", p)
}
func printFuncPointer(pp *MyPoint)  {
	(*pp).x = 1
	(*pp).y = 2
	//或
	pp.x = 1
	pp.y = 2
	fmt.Printf("-> %v\n",*pp)
}

const MAX  int = 3
func main()  {
	p := MyPoint{0, 0}
	printFunValue(p)
	pp :=&MyPoint{0,0}
	printFuncPointer(pp)

	var ptr *int
	fmt.Printf("%x\n",ptr)

	arr := []int{1, 2, 3}
	var i int
	for i = 0; i < MAX; i++{
		fmt.Printf("arr[%d] = %d\n",i, arr[i])
	}
//指向指针的指针变量
 a := 200
 var ptr1 *int
 var ptr2 **int
 /* 指针ptr1的地址 */
 ptr1 = &a
 /* 指向指针ptr1的地址 */
 ptr2 = &ptr1
 /* 获取ptr1的值 */

fmt.Printf("a = %d\n", a)
fmt.Printf("ptr1 =%d\n", *ptr1)
fmt.Printf("ptr2 =%d\n",**ptr2)







































}