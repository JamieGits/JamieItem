//当前程序的包名
package main
//导入其他的包
import (
	"fmt"
	"strconv"
)
//常量的定义
const pt  = 3.14
//全局变量的声明和赋值
var name = "gopher"
//一般类型的声明
var newType int
//结构体的声明
type gopher struct {

}
//接口的声明
type golang interface {

}

const (
	B float32 = 1 << (iota * 10)
	KB 
)
//由main函数作为程序入口启动src/com/test1.go:20
func main()  {
	var f string
	f = "中国风"
	var s int = 98
	d := strconv.Itoa(s)
	fmt.Println("hello word")
	fmt.Println(f)
	fmt.Println(d)
	fmt.Println(B)
	fmt.Println(KB)

	    a := 2
	switch {
	case a >= 0:
	fmt.Println("a = 0")
	fallthrough
	case a >= 1:
	fmt.Println("a = 1")
	default:
		fmt.Println("over")
	}
  
}





























