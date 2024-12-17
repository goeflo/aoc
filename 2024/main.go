package main

import (
	"aoc/2024/day1"
	"aoc/2024/day2"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Printf("aoc 2024 ...\n")

	var day int
	var example bool
	flag.IntVar(&day, "d", 0, "solve for this day")
	flag.BoolVar(&example, "e", false, "use example input")
	flag.Parse()

	if day == 0 {
		panic(fmt.Errorf("no day input"))
	}

	inputFilename := ""
	if example {
		inputFilename = filepath.Join(fmt.Sprintf("day%v", day), "example.txt")
	} else {
		inputFilename = filepath.Join(fmt.Sprintf("day%v", day), "puzzle.txt")
	}

	fmt.Printf("puzzle input file: %v\n", inputFilename)

	file, err := os.ReadFile(inputFilename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")

	switch day {
	case 1:
		day1.Solve(lines)
	case 2:
		day2.Solve(lines)
	default:
		fmt.Printf("day %v not yet implemented\n", day)
	}
}
