package main

import "fmt"
/**
 *  求阶乘
 */
func main()  {
	var n = 15
 fmt.Printf("%d的阶乘: %d\n",n, factorial(uint64(n)) )
 fmt.Println()
}

func factorial(n uint64) (result uint64) {
	if(n>0){
		result = n*factorial(n-1)
		return result
	}
	return 1
}