package main

import (
	"math"
	"os"
	"slices"
	"fmt"
	"strconv"
	"strings"
)

	
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func readInput() ([]int, []int) {
	data, err := os.ReadFile("input.txt")
	check(err)

	lines := strings.Split(string(data), "\n")

	var line1 []int
	var line2 []int

	for _, line := range lines {
		nums := strings.Split(line, "   ")
		num1, err := strconv.Atoi(nums[0])
		check(err)
		num2, err := strconv.Atoi(nums[1])
		check(err)
		line1 = append(line1, num1)
		line2 = append(line2, num2)
	}

	return line1, line2
}

func GetFrequency(line []int) map[int]int {
	freq := make(map[int]int)
	for _, num := range line {
		freq[num]++
	}
	return freq
}

func main() {
	line1, line2 := readInput()

	slices.Sort(line1)
	slices.Sort(line2)

	diff := 0
	for i := 0; i < len(line1); i++ {
		diff += int(math.Abs(float64(line1[i] - line2[i])))
	}

	fmt.Println(diff)

	freq2 := GetFrequency(line2)

	score := 0

	for i := 0; i < len(line1); i++ {
		data, ok := freq2[line1[i]]
		if !ok {
			data = 0
		}
		score += line1[i] * data
	}

	fmt.Println(score)
}
