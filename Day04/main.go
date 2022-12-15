package main

import (
	"adventOfCode/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	FullyContainTotal := 0
	OverlapTotal := 0
	utils.ReadFile("Day04/input.txt", func(text string) {
		ranges := strings.Split(text, ",")

		parts1 := strings.Split(ranges[0], "-")
		start1, _ := strconv.Atoi(parts1[0])
		end1, _ := strconv.Atoi(parts1[1])

		parts2 := strings.Split(ranges[1], "-")
		start2, _ := strconv.Atoi(parts2[0])
		end2, _ := strconv.Atoi(parts2[1])

		if start1 <= end2 && start2 <= end1 {
			OverlapTotal++
			if (start1 <= start2 && end1 >= end2) || (start2 <= start1 && end2 >= end1) {
				FullyContainTotal++
			}
		}

	})
	fmt.Println("Total Fully Contains:", FullyContainTotal)
	fmt.Println("Total Overlapping:", OverlapTotal)

}
