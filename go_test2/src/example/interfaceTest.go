package main

import "fmt"

type Phone interface {
	call()
}

type NikoPhone struct {
	
}

func (nikoPhone NikoPhone) call() {
	fmt.Println("I am niko,I can call you")
}

type JackPhone struct {

}

func (jackPhone JackPhone) call() {
	fmt.Println("I am jack,I can call you")
}
func main()  {
	var phone Phone
	phone = new (NikoPhone)
	phone.call()
	phone = new (JackPhone)
	phone.call()
}
//上面例子中，定义了一个接口Phone，里面有个call()方法，main中定义 了一个Phone类型变量，分别赋值NikoPhone和JackPhone，然后调用call()