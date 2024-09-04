package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("Please, provide a file name of file that should be read.")
	}

	fileName := os.Args[1]

	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		panic("There were issues trying to read the file.")
	}

	fmt.Println(string(fileContent))
}
