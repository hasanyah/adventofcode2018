package main

import (
    "bufio"
    "flag"
    "fmt"
    "os"
    "strings"
    "strconv"
)

type Claim struct {
    Id int
    X int
    Y int
    Width int
    Height int
    Safe bool
}

type Claimee struct {
    Claimer int
    Claimed int
}

var inputFile = flag.String("inputFile", "day03.input", "Relative file path to use as input.")

func main() {
    flag.Parse()
    file, err := os.Open(*inputFile)
    if err != nil {
        fmt.Printf("Had an error while opening the file %s: %v\n", *inputFile, err)
        return
    }
    defer file.Close()

    reader := bufio.NewReader(file)

    mMap := make(map[int]map[int]*Claimee)
    objMap := make(map[int]*Claim)
    for {
        line, err := reader.ReadString('\n')
        if err != nil || len(line) == 0 {
            break
        }
        line = line[:len(line)-2]
        data := strings.FieldsFunc(line, Split)
        
        id, err := strconv.Atoi(strings.Trim(data[0], " "))
        if err != nil {
            fmt.Printf("Can't parse %v: %v", data[0], err)
            break
        }
        x, err := strconv.Atoi(strings.Trim(data[1], " "))
        if err != nil {
            fmt.Printf("Can't parse %v: %v", data[1], err)
            break
        }
        y, err := strconv.Atoi(strings.Trim(data[2], " "))
        if err != nil {
            fmt.Printf("Can't parse %v: %v", data[2], err)
            break
        }
        width, err := strconv.Atoi(strings.Trim(data[3], " "))
        if err != nil {
            fmt.Printf("Can't parse %v: %v", data[3], err)
            break
        }
        height, err := strconv.Atoi(strings.Trim(data[4], " "))
        if err != nil {
            fmt.Printf("Can't parse %v: %v", data[4], err)
            break
        }
        obj := Claim{id, x, y, width, height, true}
        objMap[id] = &obj
        for i := x; i < width+x; i++ {
            if mMap[i] == nil {
                mMap[i] = make(map[int]*Claimee)
            }

            for j := y; j < height+y; j++ {
                if mMap[i][j] == nil {
                    mMap[i][j] = &Claimee{id, 1}
                } else {
                    objMap[mMap[i][j].Claimer].Safe = false
                    objMap[id].Safe = false
                    mMap[i][j].Claimed++
                }
            }
        }
    }

    result := 0

    for _, v := range mMap {
        for _, v2 := range v {
            if v2.Claimed > 1 {
                result++
            }
        }
    }
    for _, v := range objMap {
        if v.Safe {
            fmt.Printf("Winner Loner Zoner: %v\n", v)
        }
    }

    fmt.Printf("Result is: %d\n", result)
}

func Split(r rune) bool {
    return r == '#' || r == '@' || r == ',' || r == ':' || r == 'x'
}
