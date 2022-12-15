package main

import (
	"adventOfCode/utils"
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var (
		NumberList []int64
		maxSum     int64
		currentSum int64
	)

	maxSum = 0
	currentSum = 0

	utils.ReadFile("Day01/Day1.txt", func(text string) {
		if text == "" {
			NumberList = append(NumberList, currentSum)
			currentSum = 0
			return
		}

		intNumber, err := strconv.ParseInt(text, 10, 64)
		if err != nil {
			panic(err)
		}

		currentSum += intNumber
	})

	sortList(NumberList)

	for i := 0; i < 3; i++ {
		maxSum += NumberList[i]
	}

	fmt.Println("Maximum sum:", maxSum)
}

func sortList(numList []int64) {
	more := func(i, j int) bool { return numList[i] > numList[j] }
	sort.Slice(numList, more)
}
