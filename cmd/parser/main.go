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
	result:=processing.GetLogLevel(message)
	if result==processing.LevelCrit {
		fmt.Println("Critical Error Detected")
		os.Exit(1)
	}else{
		fmt.Println("Log valid")
	}
    
    // fmt.Println(result)
}