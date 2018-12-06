package main

import (
    "bufio"
    "flag"
    "fmt"
    "math"
    "os"
    "strconv"
    "strings"
)

type poi struct {
    safe bool
    area int
    posx int
    posy int
}

type zone struct {
    thepoi int
    totaldistance float64
}

var inputFile = flag.String("inputFile", "day06.input", "Relative file path to use as input.")

func main() {
    flag.Parse()
    file, err := os.Open(*inputFile)
    if err != nil {
        fmt.Printf("Had an error while opening the file %s: %v\n", *inputFile, err)
        return
    }
    defer file.Close()

    reader := bufio.NewReader(file)
    zones := make(map[int]map[int]*zone)
    pois := make(map[int]*poi)
    currPoi := 0
    maxX := 0
    maxY := 0
    for {
        line, err := reader.ReadString('\n')
        if err != nil || len(line) == 0 {
            break
        }
        line = line[:len(line)-2]
        coor := strings.Split(line, ",")
        x, err := strconv.Atoi(strings.Trim(coor[0], " "))
        if err != nil {
            fmt.Printf("Error getting x: %v\n", err)
        }
        y, err := strconv.Atoi(strings.Trim(coor[1], " "))
        if err != nil {
            fmt.Printf("Error getting y: %v\n", err)
        }
        if x > maxX {
            maxX = x
        }
        if y > maxY {
            maxY = y
        }
        pois[currPoi] = &poi{safe: true, area: 0, posx: x, posy: y}

        currPoi++
    }
    maxX++
    maxY++
    fmt.Printf("MaxX: %d, MaxY: %d\n", maxX, maxY)
    for i := 0; i < maxX; i++ {
        zones[i] = make(map[int]*zone)
        for j := 0; j < maxY; j++ {
            zones[i][j] = &zone{thepoi: 0, totaldistance: 0.00}
            distance := 10000.00
            poipos := 0
            enemy := false
            for k, v := range pois {
                newdistance := math.Abs(float64(v.posx - i)) + math.Abs(float64(v.posy - j))
                zones[i][j].totaldistance += newdistance

                if newdistance < float64(distance) {
                    poipos = k
                    distance = newdistance
                    enemy = false
                } else if newdistance == distance {
                    enemy = true
                }
            }
            if !enemy {
                pois[poipos].area++
                zones[i][j].thepoi = poipos
                if i == 0 || j == 0 || i == maxX-1 || j == maxY-1 {
                    pois[poipos].safe = false
                }
            }

        }
    }
    result := 0
    for k, v := range pois {
        fmt.Printf("Zone: %t, %d, Area: %d\n", v.safe, k, v.area)
        if v.safe && v.area > result {
            result = v.area
        }
    }
    safezone := 0
    for i := 0; i < maxX; i++ {
        for j := 0; j < maxY; j++ {
            if zones[i][j].totaldistance < 10000.00 {
                safezone++
            }
        }
    }


    fmt.Printf("Result is: %d\n", result)
    fmt.Printf("Safe zone count: %d\n", safezone)
}