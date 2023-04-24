package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)

	// for {
	// 	fmt.Println("-> ")

	// 	// readString method returns 2 things string -> input & error
	// 	// _ means that i don't care about the error
	// 	userInput, _ := reader.ReadString('\n')
	// 	userInput = strings.Replace(userInput, "\n", "", -1)

	// 	// searching for enter and replacing with empty string
	// 	// -1 means for all places

	// 	if userInput == "quit" {
	// 		break
	// 	} else {
	// 		fmt.Println(userInput)
	// 	}

	// }

	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	// fmt.Println("Press any key on keyboard. Press ESC to quit.")
	fmt.Println("MENU")
	fmt.Println("----")

	coffees := make(map[int]string)
	coffees[1] = "Cappucino"
	coffees[2] = "Latte"
	coffees[3] = "Americano"
	coffees[4] = "Mocha"
	coffees[5] = "Macchiato"
	coffees[6] = "Expresso"

	fmt.Println("1 -> Cappucino")
	fmt.Println("2 -> Latte")
	fmt.Println("3 -> Americano")
	fmt.Println("4 -> Mocha")
	fmt.Println("5 -> Macchiato")
	fmt.Println("6 -> Espresso")
	fmt.Println("Q -> Quit")

	for {
		char, _, err := keyboard.GetSingleKey()

		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(string(char))
		// t := fmt.Sprintf("You choose ", string(char))

		// if key != 0 {
		// 	fmt.Println("You pressed", char, key)
		// } else {
		// 	fmt.Println("You pressed", char)
		// }

		if char == 'q' || char == 'Q' {
			break
		} else if char != '1' && char != '2' && char != '3' && char != '4' && char != '5' && char != '6' {
			fmt.Println("Worng Option!!!")
			break
		}
		// string(char) converts '1' to 1 of type string

		i, _ := strconv.Atoi(string(char)) // convert alpha numberic to int i.e 1 of type string to 1 of type int
		fmt.Println("You choose ", coffees[i])
	}

	fmt.Println("Program exiting...")
}
