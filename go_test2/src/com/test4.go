package main

import "fmt"


func toAsPage(p *int32) *int32 {
	if p == nil {
		return nil
	}
	ret := *p + 1
	return &ret
}

func main()  {
	var a  int32 = 4
	var ptr  *int32
	ptr = &a
	//fmt.Print(*ptr)
	fmt.Println(toAsPage(ptr))
}