package main

import (
    "bufio"
    "flag"
    "fmt"
    "os"
    "regexp"
    "sort"
)

type Step struct {
    Name string
    Requirements []string
    Duration int
}

var inputFile = flag.String("inputFile", "day07.input", "Relative file path to use as input.")
var r = regexp.MustCompile("Step ([A-Z+]) must be finished before step ([A-Z+]) can begin.")
var parallelWorkers = 5
var handicap = 60

func main() {
    flag.Parse()
    file, err := os.Open(*inputFile)
    if err != nil {
        fmt.Printf("Had an error while opening the file %s: %v\n", *inputFile, err)
        return
    }
    defer file.Close()

    steps := make(map[string]*Step)
    reader := bufio.NewReader(file)
    for {
        line, err := reader.ReadString('\n')
        if err != nil || len(line) == 0 {
            break
        }
        line = line[:len(line)-1]
        parsed := r.FindStringSubmatch(line)
        if len(parsed) == 0 {
            fmt.Println("Unable to parse line.")
            continue
        }
        left := parsed[1]
        right := parsed[2]
        
        if _, ok := steps[left]; !ok {
            steps[left] = &Step{left, []string{}, Duration(left)+handicap}
        }     
        
        if _, ok := steps[right]; !ok {
            steps[right] = &Step{right, []string{left}, Duration(right)+handicap}
        } else {
            steps[right].Requirements = append(steps[right].Requirements, left)
        }
    }
    Best(steps, []string{}, 0, 0)
}

func Best(steps map[string]*Step, ongoing []string, spentTime int, forwardTime int) {

    fmt.Printf("Step size %d, ongoing size %d, spent time %d, forward time %d\n", len(steps), len(ongoing), spentTime, forwardTime)
    if len(steps) == 0 {
        fmt.Printf("No step left, leaving!!! Spent time: %d\n", spentTime)
        return
    }
    availableSteps := []string{}
    toBeDeleted := []string{}
    
    fmt.Println("Checking steps for completed ones")
    for _, v := range ongoing {
        steps[v].Duration -= forwardTime
        if steps[v].Duration == 0 {
            toBeDeleted = append(toBeDeleted, v)
        }
    }

    fmt.Println("Removing steps from referenced places")
    for _, v := range toBeDeleted {
        fmt.Printf("Deleting %s\n", v)
        for _, v2 := range steps {
            for k3, v3 := range v2.Requirements {
                if v3 == v {
                    fmt.Printf("Removing %s from requirements of %s\n", v, v2.Name)
                    v2.Requirements = append(v2.Requirements[:k3], v2.Requirements[k3+1:]...)
                    break
                }
            }
        }
        for k2, v2 := range ongoing {
            if v == v2 {
                fmt.Printf("Removing %s from ongoing tasks\n", v)
                ongoing = append(ongoing[:k2], ongoing[k2+1:]...)
                delete(steps, v)
                break
            }
        }
    }
    spentTime += forwardTime
    fmt.Printf("Spent time: %d\n", spentTime)

    for k, v := range steps {
        if len(v.Requirements) == 0 {
            fmt.Printf("%s is available\n", k)
            availableSteps = append(availableSteps, k)
        }
    }

    fmt.Printf("Sorting available list. Size: %d\n", len(availableSteps))
    sort.Strings(availableSteps)

    checked := false
    for len(ongoing) < parallelWorkers && !checked && len(availableSteps) > 0 {
        fmt.Printf("Ongoing list size: %d, parallel worker size: %d\n", len(ongoing), parallelWorkers)
        for _, v := range availableSteps {
            found := false
            for _, v2 := range ongoing {
                if v2 == v {
                    found = true
                }
            }
            if !found && len(ongoing) < parallelWorkers{
                fmt.Printf("%s will start\n", v)
                ongoing = append(ongoing, v)
            }
            checked = true
        }
    }

    minTime := 999
    for _, v := range ongoing {
        if steps[v].Duration < minTime {
            minTime = steps[v].Duration
        }
    }
    if minTime == 999 {
        minTime = 0
    }
    fmt.Printf("Work time: %d\n", minTime)
    Best(steps, ongoing, spentTime, minTime)
}

func Duration(s string) int {
    switch s {
    case "A":
        return 1
    case "B":
        return 2
    case "C":
        return 3
    case "D":
        return 4
    case "E":
        return 5
    case "F":
        return 6
    case "G":
        return 7
    case "H":
        return 8
    case "I":
        return 9
    case "J":
        return 10
    case "K":
        return 11
    case "L":
        return 12
    case "M":
        return 13
    case "N":
        return 14
    case "O":
        return 15
    case "P":
        return 16
    case "Q":
        return 17
    case "R":
        return 18
    case "S":
        return 19
    case "T":
        return 20
    case "U":
        return 21
    case "V":
        return 22
    case "W":
        return 23
    case "X":
        return 24
    case "Y":
        return 25
    case "Z":
        return 26
    default:
        return 0
    }
}
