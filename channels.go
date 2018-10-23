package main

import "fmt"

type nums struct {
	num1, num2, num3, num4 int
}

func main()  {
	messages := make(chan string)

	numbers := make(chan nums)

	go func() { numbers <- nums{1,2,3,4}}()

	go func() { messages <- "ping"}()

	msg := <-messages
	nms := <-numbers
	fmt.Println(msg)
	fmt.Println(nms)
}
