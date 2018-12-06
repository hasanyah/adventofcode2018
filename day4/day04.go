package main

import (
    "bufio"
    "flag"
    "fmt"
    "os"
    "strings"
    "strconv"
)

var inputFile = flag.String("inputFile", "day04.input", "Relative file path to use as input.")

func main() {
    flag.Parse()
    file, err := os.Open(*inputFile)
    if err != nil {
        fmt.Printf("Had an error while opening the file %s: %v\n", *inputFile, err)
        return
    }
    defer file.Close()

    reader := bufio.NewReader(file)
    for {
        line, err := reader.ReadString('\n')
        if err != nil || len(line) == 0 {
            break
        }
        line = line[:len(line)-2]
    }

    result := 0
    fmt.Printf("Result is: %d\n", result)
}
