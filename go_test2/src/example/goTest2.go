package main
//指定变量类型，声明后不赋值，使用默认值
var x,y int
 var(
 	a int
 	b bool

 )

 var str string="钟无艳"
 var c,d  =1,2
 //根据值自行判断变量类型
 var e,f =123,"hello"
func main()  {
	//省略var，注意:=左侧的变量不应该是已经声明过的，否则会编译报错(只能出现在函数体中)
	g,h:=123,"hello"
	g=110
	println(x,y,a,b,c,d,e,f,g,h,str)
	println(&c)
	//println(g)
}