package main

import (
	"fmt"
	"log"
	"myapp5/interfaces"
	"sort"
)

// basic -> numbers, strings, booleans
// aggregate -> arrays, structs
// reference -> pointers, slices, maps, functions, channels
// interface

var myInt int
var myUint uint

var myFloat32 float32
var myFloat64 float64

var myString string

var myBool bool

type myStruct struct {
	name string
	age  int
}

func main() {

	// basic
	// basic()

	// arrays
	// arrays()

	// structs
	// structs()

	// pointers()
	// slices()
	// maps()

	// functions.Main()

	// channels.Main()

	interfaces.Main()
}

func basic() {
	log.Println(myInt, myUint, myFloat32, myFloat64, myString, myBool)
}

func arrays() {
	var arr [3]string
	arr[0] = "cat"
	arr[1] = "dog"
	arr[2] = "cow"

	fmt.Printf("%s , %s, %s\n", arr[0], arr[1], arr[2])
}

func structs() {
	x := myStruct{
		name: "Vrinda",
		age:  24,
	}
	fmt.Printf("You are %s and your age is %d\n", x.name, x.age)
}

func pointers() {
	x := 10

	myPointer := &x

	fmt.Println("x is ", x)
	fmt.Println("myPointer is ", myPointer)

	*myPointer = 15
	fmt.Println("Now x is ", x)
	changeValue(&x)
	fmt.Println("Afer function call, now x is ", x)
}

func changeValue(pointer *int) {
	*pointer = 25
}

func slices() {
	var animals []string
	animals = append(animals, "dog")
	animals = append(animals, "cat")
	animals = append(animals, "horse")
	animals = append(animals, "cow")

	fmt.Println(animals)
	for i, x := range animals {
		fmt.Println(i, x)
	}
	for i := range animals {
		fmt.Println(animals[len(animals)-i-1])
	}

	fmt.Println("At 0 is", animals[0])
	fmt.Println("From 0 to 2 are", animals[0:2])
	fmt.Println("Length is", len(animals))
	fmt.Println("Is it sorted?", sort.StringsAreSorted(animals))
	sort.Strings(animals)
	fmt.Println("How about now?", animals)
	fmt.Println("Order is", animals)

	animals = DeleteFromSlice(animals, 1)
	fmt.Println("After deletion ->", animals)
}

func DeleteFromSlice(arr []string, i int) []string {
	arr[i] = arr[len(arr)-1]
	arr[len(arr)-1] = ""
	arr = arr[0 : len(arr)-1]
	return arr
}

func maps() {
	intMap := make(map[string]int)
	intMap["one"] = 1
	intMap["two"] = 2
	intMap["three"] = 3
	intMap["four"] = 4
	intMap["five"] = 5

	for key, val := range intMap {
		fmt.Println(key, val)
	}

	// delete(intMap, "four")
	val, ok := intMap["four"]
	if ok {
		fmt.Println(val, "is present.")
	} else {
		fmt.Println(val, "is not present.")
	}

	intMap["four"] = 8
}
