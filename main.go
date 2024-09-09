package main

import (
	"Gman/parsers"
	"fmt"
	"os"
)

const argsLength = 2
func main() {
	if len(os.Args) < argsLength {
		fmt.Println("Usage: go run main.go <filename>")
		return 
	}

	filename := os.Args[1]
	// need to pass the file parser the command context
	commandContext := parsers.CreateCommandContext()

	parser := parsers.CreateFileParser(filename, &commandContext)
	err :=parser.ParseFile()
	if err != nil{
		fmt.Printf("we go a problem here %s",err.Error())
	}
}
