package main

import "fmt"

func main()  {
	var a = 10
	var  ip *int
	ip = &a  //指针变量的存储地址
	fmt.Printf("变量的地址: %x\n", &a)  //变量地址十六进制
	fmt.Printf("ip变量存储的指针地址: %x\n", ip)
    /* 使用指针访问值 */
    fmt.Printf("*ip变量的值: %d\n", *ip)
    var ptr *int
    /* 空指针 */
    fmt.Printf("ptr的值为: %x\n", ptr)

}
