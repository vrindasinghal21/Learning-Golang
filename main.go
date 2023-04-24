package main

import "fmt"

func main() {
	// fmt.Println("Hello World!")

	// var x string       // 1st method
	// x = "hell World!"

	x := "hell World!" // 2nd method
	printHello(x)
}

func printHello(input string) {
	fmt.Println(input)
}
