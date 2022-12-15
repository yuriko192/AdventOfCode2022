package main

import (
	"adventOfCode/utils"
	"fmt"
	"strings"
)

func main() {
	utils.ReadFile("Day06/input.txt", func(text string) {
		text = strings.TrimSpace(text)
		numOfChar := 14
		for i := 0; i <= len(text)-numOfChar; i++ {
			sequence := text[i : i+numOfChar]
			if isUnique(sequence) {
				fmt.Println(i+numOfChar, sequence)
			}
		}
	})
}

func isUnique(sequence string) bool {
	counts := map[rune]int{}
	for _, r := range sequence {
		if _, ok := counts[r]; ok {
			return false
		}
		counts[r]++
	}
	return true
}
