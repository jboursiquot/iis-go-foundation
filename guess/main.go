package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	src := rand.NewSource(time.Now().UnixNano())
	winningNum := rand.New(src).Intn(100)
	maxGuesses := 5
	guesses := 0

	for {
		if guesses == maxGuesses {
			fmt.Println("You've run out of guesses.")
			break
		}
		guesses++

		fmt.Print("Guess a number between 0 and 100: ")

		var guessedNum int
		fmt.Scanln(&guessedNum)

		if guessedNum == winningNum {
			fmt.Println("You win!")
			break
		}

		if guessedNum < winningNum {
			fmt.Println("Higher")
			continue
		}

		if guessedNum > winningNum {
			fmt.Println("Lower")
			continue
		}
	}

	fmt.Printf("Number: %d\nGuesses: %d\nBye\n", winningNum, guesses)
}
