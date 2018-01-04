package main

import "fmt"

func main()  {
	numbers := []int{0,1,2,3,4,5,6,7,8}

	printSlice(numbers)
	/* 打印原始切片 */
	fmt.Println("numbers ==",numbers)

	fmt.Println("numbers[1:4] ==",numbers[1:4])

	fmt.Println("numbers[:3] ==",numbers[:3])

	fmt.Println("numbers[4:] ==",numbers[4:])

	fmt.Println("numbers[:] ==",numbers[:])

	number1 := numbers[2:8]
	printSlice(number1)

	number2 := numbers[:2]
	printSlice(number2)
}