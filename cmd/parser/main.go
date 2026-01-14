package main

import (
    "fmt"
    "github.com/Sumit031202/log-parser/internal/processing"
)

func main() {
    message := "Error: Database connection failed"
    
    // Calling the function from our internal package
    result := processing.GetLogLevel(message)
    
    fmt.Println(result)
}