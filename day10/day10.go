package main

import (
    "bufio"
    "flag"
    "fmt"
    "os"
    "regexp"
    "strconv"
    "strings"
)

var inputFile = flag.String("inputFile", "day10.input", "Relative file path to use as input.")
var r = regexp.MustCompile("position=<\\s?(.+),\\s?(.+)> velocity=<\\s?(.+),\\s?(.+)>")
var xdim = 100
var ydim = 100
var seconds = 10346

func main() {
    flag.Parse()
    file, err := os.Open(*inputFile)
    if err != nil {
        fmt.Printf("Had an error while opening the file %s: %v\n", *inputFile, err)
        return
    }
    defer file.Close()
    points := make(map[int]map[int]int)
    timex := make(map[int]map[int]int)
    timey := make(map[int]map[int]int)
    reader := bufio.NewReader(file)
    count := 0
    maxCommonLine := 0
    maxCommonColumn := 0
    // maxAtTime := 0
    for {
        line, err := reader.ReadString('\n')
        if err != nil || len(line) == 0 {
            break
        }
        line = line[:len(line)-2]
        parsed := r.FindStringSubmatch(line)
        if len(parsed) == 0 {
            fmt.Println("Unable to parse line.")
            continue
        }
        posx, _ := strconv.Atoi(strings.Trim(parsed[1], " "))
        posy, _ := strconv.Atoi(strings.Trim(parsed[2], " "))
        velx, _ := strconv.Atoi(strings.Trim(parsed[3], " "))
        vely, _ := strconv.Atoi(strings.Trim(parsed[4], " "))
        points[count] = make(map[int]int)
        points[count][0] = posx
        points[count][1] = posy
        points[count][2] = velx
        points[count][3] = vely
        count++
    }

    for i := 0; i < seconds; i++ {
        timex[i] = make(map[int]int)
        timey[i] = make(map[int]int)
        canvas := [100][100]string{}
        for k, v := range canvas {
            for k2, _ := range v {
                canvas[k][k2] = "."
            }
        }
        if i > 10344 {
            for _, v := range points {
                x := v[0] + i * v[2] - 150
                y := v[1] + i * v[3] - 100
                timex[i][x]++
                timey[i][y]++
                if x >= 0 && y >= 0 && x < 150 && y < 150{
                    canvas[y][x] = "#"
                    fmt.Printf("Locations: %d %d \n", x, y)
                }
            }
        }
        for _,v := range timex {
            for _,v2 := range v {
                if v2 > maxCommonColumn {
                    maxCommonColumn = v2
                    // maxAtTime = i
                }
            }
        }
        for _,v := range timey {
            for _,v2 := range v {
                if v2 > maxCommonLine {
                    maxCommonLine = v2
                }
            }
        }
        // if i > 10340 || i % 500 == 0{
        //     fmt.Printf("Max Common Line and Column: %d, %d, seconds at %d\n", maxCommonLine, maxCommonColumn, maxAtTime)
        // }

        if i == 10345 {
            for _,v := range canvas {
                fmt.Println(v)
            }
        } 
    }



}
