package main

import "fmt"

func intSeq() func(int) int {
	i := 0
	return func(int c) int {
		return i + c
	}
}

func main()  {

	nextInt := intSeq()
	ab := intSeq()

	foo, _ := fmt.Println(nextInt(5))
	fmt.Println(foo)

	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInt := intSeq()
	fmt.Println(newInt())
}