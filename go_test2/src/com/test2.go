package main

import "fmt"

func main()  {
	//break
lable1:
	for{
		for i := 0; i<10 ; i++ {
			if i>3 {
				break lable1
			}
		}
	}
	fmt.Println("ok1")

	//goto
	for{
		for i := 0; i<10; i++ {
			if i>3 {
				goto lable2
			}
		}
	}
	lable2 : fmt.Println("ok2")

	//continue
	lable3 :
		for i := 0; i<10; i++ {
			for{
				continue lable3
			}

		}
		fmt.Println("ok3")

		x, y := 1, 2
		a := []*int{&x, &y}
		fmt.Println(a)

		b := [2]int{1,2}
		c := [2]int{1,3}
		fmt.Println(b == c)
		//使用new关键字创建数组，返回指向数组的指针
		p := new([2]int)
		fmt.Println(p)

		d := [2][3]int{{1,2,3},{4,5,6}}
		fmt.Println(d)

}




























