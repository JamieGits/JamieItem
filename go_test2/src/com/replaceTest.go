package main

import "fmt"
import "strings"

func main()  {
	str := "welcome to sharejs.com"
	str = strings.Replace(str," ","",-1)
	fmt.Println(str)
	packageIds := []string{}
	d := []string{"Welcome", "for", "Tianjin", "Have", "a", "good", "journey"}
	insertSlice := []string{"It", "is", "a", "big", "city"}
	insertSliceIndex := 3
	d = append(d[:insertSliceIndex],
		append(insertSlice, d[insertSliceIndex:]...)...)

	fmt.Printf("result:%v\n", d)
	packageIds = append(packageIds,"123")
fmt.Printf("%v\n",packageIds,)

}

