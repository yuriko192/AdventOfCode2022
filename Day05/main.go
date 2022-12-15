package main

import (
	"adventOfCode/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	initialArr := []string{
		"",
		"QFMRLWCV",
		"DQL",
		"PSRGWCNB",
		"LCDHBQG",
		"VGLFZS",
		"DGNP",
		"DZPVFCW",
		"CPDMS",
		"ZNWTVMPC",
	}
	BoxesList := make([][]rune, len(initialArr))
	for i, str := range initialArr {
		BoxesList[i] = []rune(str)
	}

	utils.ReadFile("Day05/input.txt", func(text string) {
		text = strings.TrimSpace(text)

		actionList := strings.Split(text, " ")
		move_num, _ := strconv.Atoi(actionList[1])
		from, _ := strconv.Atoi(actionList[3])
		to, _ := strconv.Atoi(actionList[5])

		fromStack := BoxesList[from]
		toStack := BoxesList[to]
		defer func() {
			BoxesList[from] = fromStack
			BoxesList[to] = toStack
			fmt.Println(text)
			PrettyPrint(BoxesList)
		}()

		//for i := 0; i < move_num; i++ {
		//	movedElem := fromStack[len(fromStack)-1]
		//	fromStack = fromStack[:len(fromStack)-1]
		//	toStack = append(toStack, movedElem)
		//}

		for i := move_num; i > 0; i-- {
			movedElem := fromStack[len(fromStack)-i]
			toStack = append(toStack, movedElem)
		}
		fromStack = fromStack[:len(fromStack)-move_num]

	})
}

func PrettyPrint(BoxesList [][]rune) {
	for i, runes := range BoxesList {
		fmt.Println(i, string(runes))
	}
	fmt.Println()
}
