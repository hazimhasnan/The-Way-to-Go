package main

import (
	"fmt"
)

type Person struct {
	name  string
	age   int
	hobby string
}

//Change the hobby because of trauma
func inTrauma(p *Person) {
	p.hobby = "Dying"
}

//Change the age
func getOlder(p *Person) {
	(*p).age++
}

//Change onlt the name
func changeName(p *Person) {
	p.name = "Hazim"
}

//Change every field in the struct
func changeEverything(p *Person) {
	p.hobby = "Dying"
	(*p).age = 200
	(*p).name = "Bunny"
}

func main() {
	syaz := Person{
		name:  "syazwan",
		age:   34,
		hobby: "Sleeping",
	}

	ali := new(Person)
	ali.name = "ali"
	ali.age = 21
	ali.hobby = "meow meow"

	fmt.Printf("Value before %v, %v\n", syaz, ali)

	inTrauma(&syaz)
	getOlder(&syaz)
	changeName(&syaz)
	changeEverything(ali)

	fmt.Println(syaz, ali)
}
