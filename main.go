package main

import (
	"os"
)

func main() {
	//read command line arguments
	args := os.Args
	//check if there are enough arguments
	if len(args) < 2 {
		//print usage
		println("No input file specified")
		return
	}
	//get the input file name
	input := args[1]

}
