package packageone

import "fmt"

var privateVar = "I am private"
var PublicVar = "I am public"

func nonExported() {
	fmt.Println("I am Private method")
}

func Exported() {
	fmt.Println("I am public method")
}
