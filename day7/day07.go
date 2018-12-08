package main

import (
    "bufio"
    "flag"
    "fmt"
    "os"
    "regexp"
)

var inputFile = flag.String("inputFile", "day07.input", "Relative file path to use as input.")
var partB = flag.Bool("partB", false, "Using part b")
var r = regexp.MustCompile("Step ([A-Z+]) must be finished before step ([A-Z+]) can begin.")

func main() {
    flag.Parse()
    file, err := os.Open(*inputFile)
    if err != nil {
        fmt.Printf("Had an error while opening the file %s: %v\n", *inputFile, err)
        return
    }
    defer file.Close()

    steps := make(map[string]bool)
    reader := bufio.NewReader(file)
    recs := make(map[int]map[string]string)
    count := 0
    finalText := ""
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
        left := parsed[1]
        right := parsed[2]
        
        if _, ok := steps[left]; !ok {
            steps[left] = true
        }     
        steps[right] = false
        
        recs[count] = make(map[string]string)
        recs[count][left] = right 
        count++
    }
    Best(steps, recs, finalText)

    // for k, v := range steps {
    //     if v {
    //         finalText += k

    //     }
    // }

}

func Best(steps map[string]bool, recs map[int]map[string]string, finalText string) {
    bestLetter := "ZZ"
    for k, v := range steps {
        if v && k < bestLetter {
            bestLetter = k
        }
    }
    for _, v := range recs {
        for k2, v2 := range v {
            if k2 == bestLetter {
                steps[v2] = true
            }
        }
    }
    finalText += bestLetter
    delete(steps, bestLetter)
    Best(steps, recs, finalText)

    fmt.Println(finalText)
}