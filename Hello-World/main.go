package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println(doctor.Intro())
	for {
		fmt.Print("-> ")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1) // for windows
		userInput = strings.Replace(userInput, "\n", "", -1)   // for linux/mac

		if userInput == "quit" {
			break
		}

		fmt.Println(doctor.Response(userInput))
	}

}
