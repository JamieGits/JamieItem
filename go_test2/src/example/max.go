package main

import "fmt"

/* 函数返回两个数最大值 */
func max(num1,num2 int) int {
	if (num1 > num2) {
		return num1
	} else {
		return  num2
	}

}
/* 函数返回多个值 */
func swap(x, y string) (string, string) {
	return y, x
}


func main(){
	var a int = 1
	var b int =2
	var ret int
	ret=max(a,b)
		fmt.Printf( "最大值是 : %d\n", ret )
	fmt.Println("最大值是 :", max(1,2))

	c, d :=swap("张三","李四")
	fmt.Println(c, d)
	}
