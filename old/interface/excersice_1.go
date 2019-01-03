package main

import "fmt"

type speaker interface {
	speak()
}

type english struct {
}

type chinese struct {
}

func (e english) speak() {
	fmt.Println("Hello work")
}

func (c *chinese) speak() {
	fmt.Println("你好世界")
}

func sayHello(s speaker) {
	s.speak()
}

func main() {
	// Declare a variable of the interface speaker type
	// set to its zero value.

	var s speaker

	e := english{}

	s = e

	s.speak()

	c := chinese{}

	s = &c

	s.speak()

	sayHello(e)
	sayHello(&c)

}
