package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func CheckError(e error) {
    if e != nil {
        panic(e)
    }
}

func ReadInput() [][]int {
	data, err := os.ReadFile("input.txt")
	CheckError(err)

	lines := strings.Split(string(data), "\n")

	var result [][]int

	for _, line := range lines {
		nums := strings.Split(line, " ")
		var numsInt []int
		for _, num := range nums {
			numInt, err := strconv.Atoi(num)
			CheckError(err)
			numsInt = append(numsInt, numInt)
		}
		result = append(result, numsInt)
	}

	return result
}

func CheckIncreasing(line []int) bool {
	for i := 1; i < len(line); i++ {
		if line[i] <= line[i-1] {
			return false
		}
	}
	return true
}

func CheckDecreasing(line []int) bool {
	for i := 1; i < len(line); i++ {
		if line[i] >= line[i-1] {
			return false
		}
	}
	return true
}

func CheckAdjDiff(line []int) bool {
	for i := 1; i < len(line); i++ {
		if math.Abs(float64(line[i] - line[i-1])) > 3 {
			return false
		}
	}
	return true
}

func RemoveIndex(line []int, index int) []int {
	return slices.Delete(line, index, index+1)
}

func main() {
	input := ReadInput()

	count := 0
	for _, line := range input {
		if (CheckIncreasing(line) || CheckDecreasing(line)) && CheckAdjDiff(line) {
			count++
			continue
		}

		for i := 0; i < len(line); i++ {
			removed := RemoveIndex(slices.Clone(line), i)
			if (CheckIncreasing(removed) || CheckDecreasing(removed)) && CheckAdjDiff(removed) {
				count++
				break
			}
		}
	}

	fmt.Println(count)
}
