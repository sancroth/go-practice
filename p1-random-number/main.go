package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// untyped const btw
const maxTurns = 5

func main() {

	var winMessages [3]string = [3]string{"Woohoo good job, you found it!", "YOU WON!", "Congratulations!"}
	var loseMessages [3]string = [3]string{"Better luck next time!", "YOU LOSE!", "Shame kid, try again..."}
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Zero length params. please input a positive number")
		return
	}

	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("NaN")
		return
	}

	if guess < 0 {
		fmt.Println("Pick a positive number")
		return
	}

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(guess + 1)
		fmt.Println(n)
		if n == guess {
			fmt.Println(winMessages[rand.Intn(len(winMessages))])
			return
		}
	}
	fmt.Println(loseMessages[rand.Intn(len(loseMessages))])
}
