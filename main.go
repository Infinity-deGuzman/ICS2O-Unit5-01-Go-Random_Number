package main

import (
	"fmt"

	"math/rand"

	"time"
)

func main() {
	var userNumber int
	var random int

	// input

	fmt.Print("Guess an integer from 1 to 6: ")
	fmt.Scanln(&userNumber)

	// process
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 6

	random = rand.Intn(max-min) + min

	// output
	if random == userNumber {
		print("Good job! You guessed correctly :)")
	}

	if random != userNumber {
		print("Sorry, you got it wrong. Try again :(")
	}
}
