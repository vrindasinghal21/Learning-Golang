package main

import (
	"fmt"
	"myapp1/packageOne" // folder which you are trying to import
)

var one = "One"

func main() {
	var one = "this is block level variable"

	fmt.Println(one)

	myFunc()

	newString := packageone.PublicVar
	fmt.Println("From packageone: ", newString)
	packageone.Exported()
}

func myFunc() {
	// var one = "the number one"
	fmt.Println(one)
}
