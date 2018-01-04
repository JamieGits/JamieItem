package main

import "fmt"

func main()  {
	var a int = 100
	var b int = 200

	swap(&a, &b)
	fmt.Printf("交换后的a: %d\n", a)
	fmt.Printf("交换后的b: %d\n", b)
}

func swap(a *int, b *int)  {
	var temp int
	temp = *a
	*a = *b
	*b = temp

}