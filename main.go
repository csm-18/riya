package main

import (
	"fmt"
	"os"
	"strings"
)

// main() is the entry point to the riya compiler CLI.
func main() {
	//cli args
	args := os.Args[1:]

	//riya compiler commands parsing
	if len(args) == 0 { // about
		fmt.Println(RIYA_VERSION)
		fmt.Println("Riya is a programming language and compiler")
		fmt.Println("For help:")
		fmt.Println("  riya help")
	} else if len(args) == 1 {
		switch args[0] {
		case "version": // version
			fmt.Println(RIYA_VERSION)
		case "help": // help
			fmt.Println("Riya Compiler Commands:")
			fmt.Println("1. riya                                -> print about information")
			fmt.Println("2. riya version                        -> print compiler version")
			fmt.Println("3. riya help                           -> print compiler commands")
			fmt.Println("4. riya compile <filename.riya>        -> compile .riya file to executable")
		}
	} else if len(args) == 2 {
		if args[0] == "compile" { // compile
			filename := args[1]
			if len(filename) > 5 && strings.HasSuffix(filename, ".riya") {
				fmt.Printf("Compiling %s...\n", filename)
			}
		}
	}

}

const RIYA_VERSION = "riya 0.1.0"
