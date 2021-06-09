package main

import "fmt"

//struct definition
type Doctor struct {
	pay        int
	specialty  string
	experience int
}

func main() {
	//Ways to initialised struct

	//value type struct
	var syaz Doctor
	syaz.pay = 30
	syaz.specialty = "Unemployed"
	syaz.experience = 0

	//value type struct
	ali := Doctor{
		pay:        10,
		specialty:  "Studying",
		experience: -5,
	}

	//pointer type struct
	jaz := new(Doctor)
	jaz.pay = 200
	jaz.specialty = "Freelance"
	jaz.experience = 1

	//literal type struct
	hazim := &Doctor{
		pay:        6000,
		specialty:  "Fulltime",
		experience: 23,
	}

	fmt.Println(syaz, ali, jaz, hazim)
}
