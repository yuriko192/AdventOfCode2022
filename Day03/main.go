package main

import (
	"adventOfCode/utils"
	"fmt"
)

func First() {
	TotalScore := 0
	utils.ReadFile("Day03/input.txt", func(text string) {
		fmt.Println(text)
		n := len(text) / 2
		firstHalf := text[:n]
		secondHalf := text[n:]

		//Print the split string
		fmt.Printf("First half: %s\n", firstHalf)
		fmt.Printf("Second half: %s\n", secondHalf)

		commonChars := make(map[rune]int)
		for _, char := range firstHalf {
			commonChars[char] = 1
		}
		for _, char := range secondHalf {
			if commonChars[char] == 1 {
				commonChars[char] = 2
			}
		}

		for char, count := range commonChars {
			if count > 1 {
				//fmt.Printf("'%c' is present in both strings\n", char)
				charLoc := int(char)
				var charScore int
				//A = 65, a = 97
				if charLoc > int('a') {
					charScore = int(char) - int('a') + 1
				} else {
					charScore = int(char) - int('A') + 27
				}
				TotalScore += charScore
			}
		}
	})
	fmt.Println("Total Score:", TotalScore)

}

func Second() {
	TotalScore := 0
	CurrLine := 0
	TextArr := make([]map[rune]int, 3)
	utils.ReadFile("Day03/input.txt", func(text string) {
		fmt.Println(CurrLine, text)

		charMap := make(map[rune]int)
		for _, char := range text {
			charMap[char] = 1
		}
		TextArr[CurrLine] = charMap

		if CurrLine == 2 {
			defer func() {
				CurrLine = 0
				TextArr = make([]map[rune]int, 3)
			}()

			commonChars := TextArr[0]
			for i := 1; i < 3; i++ {
				for char, _ := range TextArr[i] {
					if commonChars[char] == i {
						commonChars[char] = i + 1
					}
				}
			}

			for char, i := range commonChars {
				if i == 3 {
					fmt.Printf("'%c' is present in all strings\n", char)
					charLoc := int(char)
					var charScore int
					//A = 65, a = 97
					if charLoc > int('a') {
						charScore = int(char) - int('a') + 1
					} else {
						charScore = int(char) - int('A') + 27
					}
					TotalScore += charScore
				}
			}
		} else {
			CurrLine++
		}
	})
	fmt.Println("Total Score:", TotalScore)

}

func main() {
	Second()
}
