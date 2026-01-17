package main

import (
	"fmt"
	"os"
	"strings"

	// "os"

	// "github.com/Sumit031202/log-parser/internal/processing"
	// "golang.org/x/text/message"
	"github.com/Sumit031202/log-parser/internal/processing"
	"github.com/Sumit031202/log-parser/internal/reader"
)

func main() {
	if len(os.Args)<2 {
		os.Exit(1)
	}
	fmt.Println("Reading the server.log...")
	fileContent:=reader.ReadFile(os.Args[1])

	// array of lines of the fileContent
	lines:=strings.Split(fileContent, "\n")

	//
	errCount:=0

	for _, line:= range lines {
		if line=="" {
			continue
		}
		level:=processing.GetLogLevel(line)

		if(level==processing.LevelCrit){
			errCount++
		}
	}
	fmt.Printf("Scan complete: caught %d critical errors",errCount)
}