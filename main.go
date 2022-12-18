package main

import (
	"fmt"
	"github.com/stovak/go-composer/pkg/composer"
	"os"
)

func main() {
	// get file from terminal
	// read the whole content of file and pass it to file variable, in case of error pass it to err variable
	parsedFile := composer.New(os.Args[1])
	// print file content
	fmt.Printf("Parsed: %+v", parsedFile)
}
