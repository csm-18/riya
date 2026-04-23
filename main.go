package main

import (
	"fmt"
	"os"
)

// main() is the entry point to the riya compiler CLI.
func main() {
	//cli args
	args := os.Args[1:]

	fmt.Println(args)
}
