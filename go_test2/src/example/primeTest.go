package main

import (
	"fmt"
)

/*
 * 输出 1-100的素数
 */
func main()  {
	var C, c int
	/* 这里不写入for循环是因为for语句执行之初会将C的值变为1，当我们goto A时for语句会重新执行(不是重新一轮的循环) */
	C = 1
	A:  for C < 100 {
		C++ //C = 1不能写入for，不然这里就不能写入
		for c = 2; c < C ; c++ {
			if  C % C == 0 {
				goto A //若发现因子则不是素数
			}
		}
		fmt.Println(C,"是素数")
	}

}
