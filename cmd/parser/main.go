package main

import (
	"fmt"
	"os"

	// "github.com/Sumit031202/log-parser/internal/processing"
	// "golang.org/x/text/message"
	"github.com/Sumit031202/log-parser/internal/processing"
	"github.com/Sumit031202/log-parser/internal/reader"
)

func main() {
	if len(os.Args)<2 {
		os.Exit(1)
	}
	fileName:=os.Args[1]
	fmt.Printf("Reading the %s...\n",fileName)
	fileContent:=reader.ReadFile(fileName)
	fmt.Printf("Scan complete: caught %d critical errors",processing.CountCritical(fileContent))
}