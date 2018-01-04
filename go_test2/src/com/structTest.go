package main

import (
	"fmt"
	"strings"
)

func StructTest01Base() {

}

//定义一个struct
type Student struct {
	id      int
	name    string
	address string
	age     int
}

//struct对象属于值类型，因此需要通过函数修改其原始值的时候必须使用指针
func modifyStudent(s Student) {
	s.name = s.name + "-modify"
}
func modifyStudentByPointer(s *Student) {
	s.name = s.name + "-modify"
}

type Person struct {
	fristName string
	lastName  string
}

//使用*Person作为参数的函数
func upPerson(p *Person) {
	p.fristName = strings.ToUpper(p.fristName)
	p.lastName = strings.ToUpper(p.lastName)
}

func structTest0101() {
	//使用new创建一个Student对象，结果为指针类型
	var s *Student = new(Student)
	s.id = 101
	s.name = "Jamie"
	s.address = "真北路"
	s.age = 18

	fmt.Printf("id:%d\n", s.id)
	fmt.Printf("name:%s\n", s.name)
	fmt.Printf("address:%s\n", s.address)
	fmt.Printf("age:%d\n", s.age)
	modifyStudentByPointer(s)
	fmt.Println(s)
}

//创建Student的其他方式
func structTest0102() {
	//使用&T{...}创建struct。结果为指针类型
	var s1 *Student = &Student{101, "John", "嘉利坊", 19}
	fmt.Println(s1)
	fmt.Println("modifyStudentByPointer...")
	modifyStudentByPointer(s1)
	fmt.Println(s1)

	//使用T{...}创建struct，结果为value类型
	var s2 Student = Student{103, "rlex", "天地港", 20}
	fmt.Println(s2)
	fmt.Println("modifyStudentByPointer...")
	modifyStudent(s2)
	fmt.Println(s2)
	//创建并初始化一个struct时，一般使用上述两种方式

	//其他方式
	var s3 *Student = &Student{id: 104, name: "Lucy"}
	fmt.Printf("s3:%d, %s, %s, %d\n", s3.id, s3.name, s3.address, s3.age)
}

func structTest0103() {

	//1- struct as a value type
	var p1 Person
	p1.fristName = "will"
	p1.lastName = "smith"
	upPerson(&p1)
	fmt.Printf("s1: %s\n", p1)

	//2- struct as a pointer
	var p2 = new(Person)
	fmt.Println(p2)
	p2.fristName = "zhangsan"
	p2.lastName = "lisi"
	(*p2).lastName = "lisi" //this is also valid
	upPerson(p2)
	fmt.Printf("p2: %s\n", p2)

	//3- struct as a literal
	var p3 = &Person{"zhangsan", "lisi"}
	upPerson(p3)
	fmt.Printf("p3: %s\n", p3)
}

func main() {
	structTest0101()
	structTest0102()
	structTest0103()

}
