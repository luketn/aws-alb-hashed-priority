package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("You must specify a hostname as the first and only command line argument.")
		os.Exit(-1)
	}

	const cap = 5000
	var priority = 1

	for _, c := range os.Args[1] {
		if (priority + int(c)) > cap {
			priority = 1
		}
		priority += int(c)
	}
	fmt.Println(priority)
}
