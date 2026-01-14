package main

import (
	"fmt"
	"os"

	"github.com/Sumit031202/log-parser/internal/processing"
	// "golang.org/x/text/message"
)

func main() {
	if(len(os.Args)<2){
		fmt.Println("no input by the user")
		return
	}
    // message := "Error: Database connection failed"
	message:=os.Args[1]
    
    // Calling the function from our internal package
    // result := processing.GetLogLevel(message)
	result:=processing.ProcessLog(message)
    
    fmt.Println(result)
}