package day1

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func Solve(input []string) {

	a1 := []int{}
	a2 := []int{}

	for _, line := range input {
		split := strings.Split(line, " ")
		i1, _ := strconv.Atoi(split[0])
		i2, _ := strconv.Atoi(split[len(split)-1])
		a1 = append(a1, i1)
		a2 = append(a2, i2)
	}

	question1(a1, a2)
	question2(a1, a2)

}

func question2(a1, a2 []int) {

	result := 0
	for i := 0; i < len(a1); i++ {
		c := countOccurence(a1[i], a2)
		result += (c * a1[i])
	}
	fmt.Printf("part two: %v\n", result)
}

func countOccurence(a int, v []int) int {
	c := 0
	for i := 0; i < len(v); i++ {
		if v[i] == a {
			c++
		}
	}
	return c
}

func question1(a1, a2 []int) {
	sort.Ints(a1)
	sort.Ints(a2)
	totalDist := 0
	for i := 0; i < len(a1); i++ {
		dist := math.Abs(float64(a1[i] - a2[i]))
		totalDist += int(dist)
	}
	fmt.Printf("part one: %v\n", totalDist)
}
