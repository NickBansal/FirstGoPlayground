package main

import (
	"fmt"
	"strconv"
)

// Person object
type Person struct {
	firstName, surName, city, gender string
	age                              int
}

func (p Person) greet() string {
	return "Hello my name is " + p.firstName + " " + p.surName + ", age: " + strconv.Itoa(p.age)
}

func (p *Person) hasBirthday() {
	level := true
	if p.age >= 30 {
		level = false
	} else if p.age <= 20 {
		level = true
	}

	if level == true {
		p.age++
	} else if level == false {
		p.age--
	}
}

func main() {
	person1 := Person{"Nick", "Bansal", "Manchester", "male", 30}

	fmt.Println(person1)
	person1.firstName = "Nicky"
	person1.surName = "Bee"
	person1.hasBirthday()
	person1.hasBirthday()
	person1.hasBirthday()
	person1.hasBirthday()
	person1.hasBirthday()

	fmt.Println(person1.age)
}
