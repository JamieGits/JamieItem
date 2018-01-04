package main
import (
"fmt"
"reflect"
)
type NotknownType struct {
	s1, s2, s3 string
}
var secret interface{} = NotknownType{"Ada", "Go", "Oberon"}
func main() {
	value := reflect.ValueOf(secret)
	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, value.Field(i))
	}
	str := "XBodyContentX"
	content := str[0:2]
	fmt.Println(content)
	rs := []rune(str)
	fmt.Println(string(rs[0:2]))



}
