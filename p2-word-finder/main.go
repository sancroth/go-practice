package main

import (
	"fmt"
	"os"
	"strings"
)

const storedText = "" +
	"the lazy cat went to the lazy super market to find lazy cat milk"

func main() {
	words := strings.Fields(storedText)
	query := os.Args[1:]

	//use of labeled continue to find the word only once
queries:
	for _, q := range query {
		tmpQ := strings.ToLower(q)
		for i, w := range words {
			if tmpQ == w {
				fmt.Printf("#%d: %q\n", i+1, w)
				continue queries
			}
		}
	}
}
