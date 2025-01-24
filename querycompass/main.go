package main

import (
    "fmt"
    "os"
    "github.com/CyberCompass/querycompass"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: querycompass input.txt [output.json]")
        os.Exit(1)
    }
    querycompass.Run(os.Args[1:])
}
