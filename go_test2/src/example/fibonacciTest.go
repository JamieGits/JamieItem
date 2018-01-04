package main

import "fmt"
/**
 *  斐波那契数列
 */
func main()  {
	var i int
	for i = 0;i < 10; i++ {
		fmt.Printf("%d\t",fibonacci(i))
	}
}

func fibonacci(i int) int {
	if(i < 2){
		return i
	}else{
		return fibonacci(i-2)+fibonacci(i-1)
	}
}
