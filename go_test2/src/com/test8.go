package main

import (
	"fmt"
	"encoding/json"
)
var array6 = []string{4: "Smith", 2: "Alice"}

//输出字符串数组切片
func rangeObjPrint(array []string) {
	for i, v := range array {
		fmt.Printf("index:%d  value:%s\n", i, v)
	}
}

var flag bool
func main()  {
	fmt.Printf("array6--- type:%T \n", array6)
	rangeObjPrint(array6)
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// We need to provide a variable where the JSON
	// package can put the decoded data. This
	// `map[string]interface{}` will hold a map of strings
	// to arbitrary data types.
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
	 fmt.Println(flag)
}