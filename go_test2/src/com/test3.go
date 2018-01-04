package main

import "fmt"

func test1()  {
	var slice2 []int
	slice := []int{1,2,3,4,5,6,7,8,9}
	for _,n := range slice {
		fmt.Printf("%d\n", n)
		if n%3 == 0 {
			slice2 = append(slice2,n)
			fmt.Printf("%d\n",n)
		}

	}
	fmt.Printf("%v\n",slice2)
}
func main()  {
	slice := append([]int{1,2,3,4}, 7,10,6)
	fmt.Println(slice)

	 //第二个参数可以直接追加到第一个切片后面，要注意的是这种用法函数的参数只能接受两个slice，并且末尾要加三个点
	slice1 := append([]int{1,2,3,5}, []int{9,5,7,1}...)
	fmt.Println(slice1)
	//将字符串当做[]byte类型作为第二个参数传入
	bytes := append([]byte("hello"), "world"...)
	fmt.Println(bytes)
	//append函数返回值必须有变量接收，不然编译器会报错
	test1()

}
