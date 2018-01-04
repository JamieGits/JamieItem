package main

import "fmt"

func main()  {
	x, y := 1, 2
	var arr =[...]int{5:2}
	//数组指针
	var pf *[6]int = &arr
	//指针数组
	pfArr := [...]*int{&x, &y}
	//数组指针pf的值得到的是一个指向arr地址的一个指针
	fmt.Println(pf)
	//指针数组pfArr则是一个数组内的元素全是指针类型
	fmt.Printf("%+v\n",pfArr)
	for v, k := range pfArr {
		fmt.Printf("%+v : %+v\n",v, *k)
	}



}
func sd()  {

}