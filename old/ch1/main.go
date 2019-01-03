package main

import "fmt"

func main() {
	inc := 10
	fmt.Println("before : ", inc, &inc)
	increment(&inc)
	fmt.Println("after", inc, &inc)
}

func increment(inc *int) {
	*inc++
	fmt.Println("inc : ", *inc, &inc, inc)
}
