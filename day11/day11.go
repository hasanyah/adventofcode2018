package main

import (
    "fmt"
)

var serial = 8444

func main() {
    maxPower := 0
    size := 0
    makerx := 0
    makery := 0
    table := [301][301]int{}
    for i := 1; i <= 300; i++ {
        for j := 1; j <= 300; j++ {
            table[i][j] = coordinatePower(i, j)
        }
    }
    for i := 1; i <= 300; i++ {
        for j := 1; j <= 300; j++ {
            maxSize := 300 - i + 1

            if j > i {
                maxSize = 300 - j + 1
            }
            gridTotal := 0
            gridSize := 0
            for k := 1; k <= maxSize; k++ {
                gridTotal = gridPower(i, j, k, table)  
                gridSize = k

                if gridTotal > maxPower {
                    maxPower = gridTotal
                    makerx = i
                    makery = j
                    size = gridSize
                }
                if i == j && k == 1 {
                    fmt.Printf("calculating x, y : %d, size: %d, power: %d\n", j, k, gridTotal)
                }
            }
        }
    }
    fmt.Printf("Total Power of %d found at x: %d y: %d, size: %d\n", maxPower, makerx, makery, size)    
}

func coordinatePower(x, y int) int {
    return ((x + 10) * y + serial) * (x + 10) / 100 % 10 - 5
}

func gridPower(x, y, z int, a [301][301]int) int {
    sum := 0
    for i := x; i < x+z; i++ {
        for j := y; j < y+z; j++ {
            sum += a[i][j]
        }       
    }

    return sum
}
