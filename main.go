package main

import (
    "fmt"
    "go-reloaded/processor"
    "os"
)

func main() {
    if len(os.Args) != 3 {
        fmt.Println("Usage: go run . <input.txt> <output.txt>")
        return
    }

    content, err := os.ReadFile(os.Args[1])
    if err != nil {
        fmt.Printf("Error reading file: %v\n", err)
        return
    }

    result := processor.Process(string(content))

    if err := os.WriteFile(os.Args[2], []byte(result), 0644); err != nil {
        fmt.Printf("Error writing file: %v\n", err)
        return
    }

    fmt.Println("Done!")
}