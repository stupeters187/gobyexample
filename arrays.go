package main

import "fmt"

func main()  {

	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 187
	fmt.Println("get:", a)
	fmt.Println("set:", a[4])

	fmt.Println("length", len(a))
	fmt.Println("capacity", cap(a))

	b := [5]int{1,2,3,4,5}
	fmt.Println("decl", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++{
		for j := 0; j < 3; j++{
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D:", twoD)
}
