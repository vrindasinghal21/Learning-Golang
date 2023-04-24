package interfaces

import "fmt"

type Animal interface {
	Says() string
	HowManyLegs() int
}

type Dog struct {
	name  string
	sound string
	legs  int
}

func (d *Dog) Says() string {
	return d.sound
}
func (d *Dog) HowManyLegs() int {
	return d.legs
}

type Cat struct {
	name  string
	sound string
	legs  int
	claws bool
}

func (c *Cat) Says() string {
	return c.sound
}
func (c *Cat) HowManyLegs() int {
	return c.legs
}

func Main() {
	dog := Dog{
		name:  "dog",
		sound: "woof",
		legs:  4,
	}
	Riddle(&dog)

	cat := Cat{
		name:  "cat",
		sound: "meow",
		legs:  4,
		claws: true,
	}
	Riddle(&cat)
}

func Riddle(a Animal) {
	riddle := fmt.Sprintf("This animal says `%s`, and has %d legs. What animal is it?", a.Says(), a.HowManyLegs())
	fmt.Println(riddle)
}
