package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/eiannone/keyboard"
)

var reader *bufio.Reader

type User struct {
	Username string
	Age      int
	favNum   float64
	likePet  bool
}

func main() {
	reader = bufio.NewReader(os.Stdin)

	var user User
	user.Username = readString("What is your name?")
	user.Age = readInt("What is your age?")
	user.favNum = readFloat("What is your favourite number?")
	user.likePet = readKey("Would you like pets? (y/n)")

	fmt.Printf("Hello %s! \nYou are %d years old.\nAnd ello! your favourite number is same as mine i.e %.2f.\n", user.Username, user.Age, user.favNum)
	if user.likePet {
		fmt.Printf("You are fond of pets.\n")
	} else {
		fmt.Println("You don't like pets.")
	}
}

func prompt() {
	fmt.Print("-> ")
}

func readString(s string) string {
	for {
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)
		if userInput == "" {
			fmt.Println("Please enter a value")
		} else {
			return userInput
		}
	}

}

func readInt(s string) int {
	for {
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)
		t, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Please enter a whole number")
		} else {
			return t
		}
	}

}

func readFloat(s string) float64 {
	for {
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)
		t, err := strconv.ParseFloat(userInput, 64)
		if err != nil {
			fmt.Println("Please enter a number")
		} else {
			return t
		}
	}

}

func readKey(s string) bool {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	for {
		fmt.Println(s)
		prompt()

		char, _, _ := keyboard.GetSingleKey()
		fmt.Println(string(char))
		if char == 'n' || char == 'N' {
			return false
		} else if char == 'y' || char == 'Y' {
			return true
		} else {
			fmt.Println("Please enter a valid answer")
		}
	}

}
