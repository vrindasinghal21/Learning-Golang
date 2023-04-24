package channels

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

// channel of type rune
var keyPressChan chan rune

func Main() {
	// go doSomething(("Hello world"))

	// fmt.Println("This is a message.")
	// for {
	// 	// do nothing
	// }

	keyPressChan = make(chan rune)

	// this will run in background for listening any key
	go listeinForKeyPress()

	fmt.Println("Press any key or q to quit")
	_ = keyboard.Open()

	defer func() {
		keyboard.Close()
	}()

	for {
		char, _, _ := keyboard.GetSingleKey()
		if char == 'q' || char == 'Q' {
			break
		}

		keyPressChan <- char
	}
}

func doSomething(s string) {
	until := 0
	for {
		fmt.Println(" s is ", s)
		until = until + 1

		if until >= 5 {
			break
		}
	}

}

func listeinForKeyPress() {
	for {
		key := <-keyPressChan
		fmt.Printf("You presses %q\n", key)
	}
}
