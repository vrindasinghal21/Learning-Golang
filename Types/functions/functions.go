package functions

import "fmt"

func Main() {
	z := addTwoNum(2, 3)

	fmt.Println("Sum from normal function (2,3) ->", z)

	x := sumMany(1, 2, 3, 4, 5)
	fmt.Println("Sum from variatic function (1,2,3,4,5) ->", x)

	d := Animal{
		name:  "dog",
		sound: "woof",
		legs:  4,
	}

	fmt.Printf("%s says %s and has %d legs.\n", d.name, d.sound, d.legs)

	d.Says()
	d.howManyLegs()

	fmt.Printf("%s says %s and has %d legs.\n", d.name, d.sound, d.legs)

}

// assigning name to return type -> rarely used in GO usually for short or one line functions because it reduce the readibility
func addTwoNum(a, b int) (sum int) {
	sum = a + b
	return
}

// Variatic function
func sumMany(nums ...int) int {
	total := 0
	for _, x := range nums {
		total = total + x
	}

	return total
}

type Animal struct {
	name  string
	sound string
	legs  int
}

// assign receiver to function
func (a *Animal) Says() {
	fmt.Printf("%s says %s", a.name, a.sound)
	fmt.Println()
	a.name = "cat"
	fmt.Printf("%s says %s\n", a.name, a.sound)

}

func (a Animal) howManyLegs() {
	fmt.Printf("%s has %d legs.\n", a.name, a.legs)
	a.legs = 5
	fmt.Printf("%s has %d legs.\n", a.name, a.legs)

}
