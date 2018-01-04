package main
/**
  * 常量测试
 */
import (
	"fmt"
	"unsafe"
)
const(
  a = "abc"
  b = len(a)
  c = unsafe.Sizeof(a)
  d
  f
)
 func main() {
	 const LENGTH int = 10
	 const WIDTH int = 5
	 var area int
	 //const a,b,c=1,false,"str"  //多重赋值
	 area = LENGTH * WIDTH
	 fmt.Printf("面积为 : %d", area)
	 println()
	 //在定义常量组时，如果不提供初始值，则表示将 使用上行的表达式的值
	 println(a, b, c, d, f)
 }
