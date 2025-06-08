package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// Target is a number between 1 and 100 inclusive
	target := random.Intn(100) + 1

	// Welcome message
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I have chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")

	var guess int

	for {
		fmt.Println("Enter your guess:")
		fmt.Scan(&guess)

		// Check if the guess is correct
		if guess == target {
			fmt.Println("Congratulations! You guessed the correct number.")
			break
		} else if guess < target {
			fmt.Println("Too low. Try again!")
		} else {
			fmt.Println("Too high. Try again!")
		}
	}

}
