package main

import "fmt"

type person struct {
	name string
	age int
}

func main()  {

	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Todd"})

	anne := &person{"anne", 187}
	fmt.Println(anne)

	s := person{"Lloyd Banks", 55}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 25
	fmt.Println(sp.age)
	fmt.Println(s)
}
