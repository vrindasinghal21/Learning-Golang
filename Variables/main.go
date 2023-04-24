package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = " and don't type your number in, just press ENTER key when ready!"

func main() {

	// seed the random number - it'll change in every nano second
	rand.Seed(time.Now().UnixNano())
	// we add 2 because we want the number used to be atleast 2, an no
	// greater than 10 (multiplyingby 1 makes the gamea bit silly)

	// one way - declare, then assign (two steps)
	var x int
	x = rand.Intn(8) + 2

	// another way - declare type and name and assign value
	var y = rand.Intn(8) + 2

	// one step variable -  declare name, assign value, and let GO figure out the type
	sub := rand.Intn(8) + 2
	var answer int // will store default value
	answer = x*y - sub

	playTheGame(x, y, sub, answer)
}

func playTheGame(x, y, sub, answer int) {
	reader := bufio.NewReader(os.Stdin)
	// display a welcome/instruction
	fmt.Println("Guess the number game")
	fmt.Println("---------------------")
	fmt.Println()
	fmt.Println("Think of a number between 1 and 10", prompt)
	reader.ReadString('\n')

	// take them through the game
	// var prompt = "hello"
	fmt.Println("Multuply you number by ", x, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply the result by ", y, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now substrac ", sub, " from the result", prompt)
	reader.ReadString('\n')

	// give them the answer
	fmt.Println("The answer is ", answer)
}
