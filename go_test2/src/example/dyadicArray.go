package main

import "fmt"

func main()  {
	var arr = [5][2]int{{0,1}, {2,3}, {4,5}, {6,7}, {8,9}}
	var i, j int
	for i = 0; i < 5; i++ {
		for j = 0; j < 2; j++ {
			fmt.Printf("arr[%d][%d] = %d\n", i, j, arr[i][j])
		}
	}
}
