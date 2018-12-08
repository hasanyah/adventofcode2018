package main

import (
    "bufio"
    "flag"
    "fmt"
    "os"
    "regexp"
    "sort"
    "strconv"
)

type Record struct {
    Month, Day, Hour, Minute int
    Onduty int
    Awake bool
}

type ByTime []Record

func (t ByTime) Len() int { return len(t) }
func (t ByTime) Swap(i, j int) { t[i], t[j] = t[j], t[i] }
func (t ByTime) Less(i, j int) bool {
    if t[i].Month != t[j].Month {
        return t[i].Month < t[j].Month
    } 

    if t[i].Day != t[j].Day {
        return t[i].Day < t[j].Day
    } 

    if t[i].Hour != t[j].Hour {
        return t[i].Hour < t[j].Hour
    } 

    return t[i].Minute < t[j].Minute
}

var inputFile = flag.String("inputFile", "day04.input", "Relative file path to use as input.")
var partB = flag.Bool("partB", false, "Using part b")
var r1 = regexp.MustCompile("\\[1518-(\\d+)-(\\d+) (\\d+):(\\d+)\\] ([Ga-z#0-9 ]+)")
var r2 = regexp.MustCompile("Guard #(\\d+) begins shift")

func main() {
    flag.Parse()
    file, err := os.Open(*inputFile)
    if err != nil {
        fmt.Printf("Had an error while opening the file %s: %v\n", *inputFile, err)
        return
    }
    defer file.Close()

    recs := make([]Record, 0)
    reader := bufio.NewReader(file)
    // guards := make(map[int]*Guard)
    // moments := make(map[int]*Moment)
    for {
        line, err := reader.ReadString('\n')
        if err != nil || len(line) == 0 {
            break
        }
        line = line[:len(line)-2]
        parsed := r1.FindStringSubmatch(line)
        if len(parsed) == 0 {
            fmt.Println("Unable to parse line.")
            continue
        }
        month, err := strconv.Atoi(parsed[1])
        day, err := strconv.Atoi(parsed[2])
        hour, err := strconv.Atoi(parsed[3])
        minute, err := strconv.Atoi(parsed[4])
        text := parsed[5]
        awake := true
        num := 0
        if guardNo := r2.FindStringSubmatch(text); len(guardNo) == 2 {
            num, err = strconv.Atoi(guardNo[1])
        } else if text == "wakes up" {
            awake = true
        } else if text == "falls asleep" {
            awake = false
        } else {
            fmt.Println("Unable to parse text.")
        }
        recs = append(recs, Record{month, day, hour, minute, num, awake})
    }

    sort.Sort(ByTime(recs))

    onduty := 0
    lastAsleepMinute := 0
    schedule := make(map[int]map[int]int)
    for idx, v := range recs {
        if v.Onduty != 0 {
            if schedule[v.Onduty] == nil {
                schedule[v.Onduty] = make(map[int]int)
            }
            onduty = v.Onduty
        } else {
            recs[idx].Onduty = onduty
            if v.Awake {
                for i := lastAsleepMinute; i < v.Minute; i++ {
                    schedule[onduty][i]++
                }
            } else {
                lastAsleepMinute = v.Minute
            }
        }
    }

    highest := 0
    guardNo := 0
    chosenMinute := 0
    simpleGuard := 0
    simpleMinute := -1
    simpleCount := 0
    for guard, minutes := range schedule {
        total := 0
        mostMinutes := 0
        guardChosenMinute := -1
        for min, count := range minutes {
            total += count
            if count > mostMinutes {
                mostMinutes = count
                guardChosenMinute = min
            }
            if count > simpleCount {
                simpleCount = count
                simpleMinute = min
                simpleGuard = guard
            }
        }
        if total > highest {
            guardNo = guard
            highest = total
            chosenMinute = guardChosenMinute
        }
    }
    if !*partB {
        fmt.Printf("Result is guard %d with minutes %d: %d\n", guardNo, chosenMinute, guardNo*chosenMinute)
    } else {
        fmt.Printf("Result is guard %d with minutes %d: %d\n", simpleGuard, simpleMinute, simpleGuard*simpleMinute)
    }
}
