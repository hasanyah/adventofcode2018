package main

import (
    "bufio"
    "flag"
    "fmt"
    "os"
)

var inputFile = flag.String("inputFile", "inputs/day02.input", "Relative file path to use as input.")

func main() {
    flag.Parse()
    file, err := os.Open(*inputFile)
    if err != nil {
        fmt.Printf("Had an error while opening the file %s: %v\n", *inputFile, err)
        return
    }
    defer file.Close()

    reader := bufio.NewReader(file)
    twos, threes := 0, 0
    memo := make(map[int]map[string]bool)
    for {
        line, err := reader.ReadString('\n')
        if err != nil || len(line) == 0 {
            break
        }
        line = line[:len(line)-1]

        for pos := range line {
            if memo[pos] == nil {
                memo[pos] = make(map[string]bool)
            }
            trunc := line[0:pos] + line[pos+1:]
            if memo[pos][trunc] {
                fmt.Printf("Shared characters: %s\n", trunc)
                break
            }
            memo[pos][trunc] = true
        }

        seen := make(map[rune]int)
        for _, c := range line {
            seen[c] ++
        }
        seen3 := false
        seen2 := false
        for _, v := range seen {
            if v == 3 && !seen3 {
                threes++
                seen3 = true
            } else if v == 2 && !seen2 {
                twos++
                seen2 = true
            }
        }
    }
    result := twos * threes
    fmt.Printf("Result is: %d\n", result)
}