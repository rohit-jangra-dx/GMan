package parsers

import (
	"bufio"
	"log"
	"os"
)


type FileParser struct {
	filePath string
	commandContext *CommandContext
}

func CreateFileParser(filePath string, commandContext *CommandContext) FileParser {
	return FileParser{
		filePath: filePath,
		commandContext: commandContext,
	}
}

func (f *FileParser) ParseFile() error{

	file, err := os.Open(f.filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		commandStr := scanner.Text()
		err = f.commandContext.ExecuteCommand(commandStr)
		if err != nil{
			return err
		}
	}
	return nil
}