package main

import (
    "fmt"
)


func main() {
    players := 477
    max := 7155951
    fmt.Printf("%d Players and %d Marbles\n", players, max)
    currentPlayer := 0
    currentMarble := 0
    currentpos := 0
    marbles := []int{}
    scores := [478]int{}
    highscore := 0
    highscorer := 0

    for currentMarble < max+1 {
        if currentMarble == 0 {
            marbles = append(marbles, 0)
            currentpos = 0
        } else if currentMarble == 1 {
            marbles = append(marbles, 1)
            currentpos = 1
        } else if currentMarble == 2 {
            marbles = append(marbles, 0)
            copy(marbles[2:], marbles[1:])
            marbles[1] = 2
            currentpos = 1
        } else if currentMarble % 23 == 0 {
            if currentpos > 7 {
                scores[currentPlayer] += currentMarble + marbles[currentpos-7]
                marbles = append(marbles[:currentpos-7], marbles[currentpos-6:]...)
                currentpos -= 7
            } else {
                scores[currentPlayer] += currentMarble + marbles[currentpos-6+len(marbles)]
                marbles = append(marbles[:currentpos-6+len(marbles)], marbles[currentpos-5+len(marbles):]...)
                currentpos = currentpos-6+len(marbles)
            }
        }else if newpos := currentpos + 2; newpos > len(marbles) {
            marbles = append(marbles, 0)
            copy(marbles[2:], marbles[1:])
            marbles[1] = currentMarble
            currentpos = 1
        } else {
            marbles = append(marbles, 0)
            copy(marbles[currentpos+3:], marbles[currentpos+2:])
            currentpos += 2
            marbles[currentpos] = currentMarble
        }
        currentMarble++
        currentPlayer++
        if currentPlayer > players {
            currentPlayer -= players
        }
        // fmt.Println(marbles)
        if currentMarble % 50000 == 0 {
            fmt.Println(currentMarble)
        }

        if currentMarble == 7085100 {
            for k, v := range scores {
                if v > highscore {
                    highscore = v
                    highscorer = k
                }
            }
            fmt.Printf("Temp highscore: %d, Highscorer: %d\n", highscore, highscorer)
        }
    }
    highscore = 0
    highscorer = 0
    for k, v := range scores {
        if v > highscore {
            highscore = v
            highscorer = k
        }
    }
    fmt.Printf("Highscore: %d, By: %d\n", highscore, highscorer)
    
}
