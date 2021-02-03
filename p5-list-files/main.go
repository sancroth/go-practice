package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {

	// empty files flag
	efPtr := flag.Bool("e", false, "If -e is used then the program will print only empty files found. if -d is used then it will also print empty dirs")
	nefPtr := flag.Bool("ne", false, "If -ne is used then the program will print only non-empty files found. if -d is used then it will also print non-empty dirs")
	dirPtr := flag.Bool("d", false, "if -d is used then the program will print all the directories found")

	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		args = append(args, ".")
	}

	files, err := ioutil.ReadDir(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, f := range files {
		if *dirPtr && f.IsDir() {
			fmt.Println(f.Name())
			continue
		}

		if f.Size() > 0 && *nefPtr {
			fmt.Println(f.Name())
		}

		if f.Size() == 0 && *efPtr {
			fmt.Println(f.Name())
		}
	}
}
