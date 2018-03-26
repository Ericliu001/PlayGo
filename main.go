package main

import (
	"fmt"
)

type person struct {
	fname string
	lname string
}

func main() {
	fmt.Println("Hello")
	marry := person{
		"Marry",
		"Lauu",
	}

	speak(marry)
}

func speak(p person) {
	fmt.Println("G'day", p.fname, p.lname)
}
