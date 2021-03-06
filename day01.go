package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day01.input", "Relative file path to use as input.")
var partB = flag.Bool("partB", false, "Using part b")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	result := 0
	seen := make(map[int]bool)
outer:
	for {
		for _, line := range strings.Split(contents, "\r\n") {
			if len(line) == 0 {
				break
			}
			offset, _ := strconv.Atoi(line)
			result += offset
			if seen[result] && *partB {
				fmt.Printf("Saw %d twice.\n", result)
				break outer
			}
			seen[result] = true
		}
	}

	fmt.Printf("%d\n", result)
}
