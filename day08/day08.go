package main

import (
    "bufio"
    "flag"
    "fmt"
    "os"
    "strconv"
)

type Node struct {
    ChildrenCount int
    MetaDataCount int
    MetaData []int
    StartPosition int
    EndPosition int
}

var inputFile = flag.String("inputFile", "day08.input", "Relative file path to use as input.")

func main() {
    flag.Parse()
    file, err := os.Open(*inputFile)
    if err != nil {
        fmt.Printf("Had an error while opening the file %s: %v\n", *inputFile, err)
        return
    }
    defer file.Close()

    reader := bufio.NewReader(file)
    list := []int{}
    nodes := []Node{}
    lastNodeStartPosition := 0
    newNode := true
    newMeta := false
    count := 0
    for {
        num, err := reader.ReadString(' ')
        if err != nil || len(num) == 0 {
            break
        }
        num = num[:len(num)-1]
        numint, _ := strconv.Atoi(num)
        list = append(list, numint)
        if newNode {
            nodes = append(nodes, Node{numint, -1, []int{}, count, -1})
            lastNodeStartPosition = count
            newNode = false
            newMeta =  true

            
        }
        count++
    }
}
