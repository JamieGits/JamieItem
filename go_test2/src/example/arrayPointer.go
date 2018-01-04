package main

import "fmt"

const MAX int = 3

func main()  {
	arr := [MAX]int{10, 20, 30}
	var i int
	var ptr [MAX]*int
	for i = 0; i < MAX; i++ {
		ptr[i] = &arr[i]   /* 整数地址赋值给指针数组 */
	}
   fmt.Printf("指针ptr的值: %x\n", ptr)
	for i = 0; i < MAX; i++ {
		fmt.Printf("ptr[%d] = %d\n", i, *ptr[i])
	}
}