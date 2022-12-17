package main

import (
	"adventOfCode/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	HighestSceneScore := int64(0)

	var fullIntArray [][]int
	utils.ReadFile("Day08/test.txt", func(text string) {
		text = strings.TrimSpace(text)
		numberStrings := strings.Split(text, "")
		intArray := make([]int, len(text))
		for i, numberString := range numberStrings {
			numberInt, _ := strconv.Atoi(numberString)
			intArray[i] = numberInt
		}
		fullIntArray = append(fullIntArray, intArray)
	})

	lenY := len(fullIntArray)
	lenX := len(fullIntArray[0])

	boolArr := make([][]bool, lenY)
	scenicArr := make([][]int64, lenY)
	for i := range boolArr {
		boolArr[i] = make([]bool, lenX)
		scenicArr[i] = make([]int64, lenX)
	}

	UpArr := utils.InitArrayWithValue(lenX, -1).([]int)
	DownArr := utils.InitArrayWithValue(lenX, -1).([]int)

	lenY -= 1
	lenX -= 1
	for j, intArr := range fullIntArray {
		LeftMax := -1
		RightMax := -1
		for i, _ := range intArr {
			sceneScore := CalculateSceneScore(fullIntArray, j, i)
			scenicArr[j][i] = sceneScore
			if sceneScore > HighestSceneScore {
				HighestSceneScore = sceneScore
			}

			if intArr[i] > LeftMax {
				boolArr[j][i] = true
				LeftMax = intArr[i]
			}
			if intArr[lenX-i] > RightMax {
				boolArr[j][lenX-i] = true
				RightMax = intArr[lenX-i]
			}
			if intArr[i] > UpArr[i] {
				boolArr[j][i] = true
				UpArr[i] = intArr[i]
			}
			if fullIntArray[lenY-j][i] > DownArr[i] {
				boolArr[lenY-j][i] = true
				DownArr[i] = fullIntArray[lenY-j][i]
			}
		}
	}

	totalCount := 0
	for _, array := range boolArr {
		for _, val := range array {
			if val {
				fmt.Print("1")
				totalCount++
			} else {
				fmt.Print("0")
			}
		}
		fmt.Println()
	}

	for _, array := range scenicArr {
		for _, val := range array {
			fmt.Printf("%d,", val)
		}
		fmt.Println()
	}

	fmt.Println("Total Visible:", fmt.Sprintf("%+v", totalCount))
	fmt.Println("Highest Scene:", fmt.Sprintf("%+v", HighestSceneScore))

}

type sceneParameter struct {
	array [][]int
	i     int
	j     int
	lenY  int
	lenX  int
}

func CalculateSceneScore(array [][]int, j int, i int) int64 {
	reqParam := sceneParameter{
		array: array,
		i:     i,
		j:     j,
		lenX:  len(array[j]) - 1,
		lenY:  len(array) - 1,
	}

	left := countSceneDistanceHori(reqParam, -1)
	right := countSceneDistanceHori(reqParam, 1)
	up := countSceneDistanceVert(reqParam, -1)
	down := countSceneDistanceVert(reqParam, 1)

	return up * left * down * right
}

func countSceneDistanceHori(param sceneParameter, direction int) int64 {
	array := param.array
	i := param.i
	j := param.j
	length := param.lenX
	treeHeight := array[j][i]

	x := i + direction
	for true {
		if x < 0 || x > length || array[j][x] >= treeHeight {
			break
		}
		x += direction
	}
	if x < 0 || x > length {
		x -= direction
	}
	distance := math.Max(float64(i-x), float64(x-i))
	distance = math.Max(1, distance)
	return int64(distance)
}

func countSceneDistanceVert(param sceneParameter, direction int) int64 {
	array := param.array
	i := param.i
	j := param.j
	length := param.lenY
	treeHeight := array[j][i]

	y := j + direction
	for true {
		if y < 0 || y > length || array[y][i] >= treeHeight {
			break
		}
		y += direction
	}
	if y < 0 || y > length {
		y -= direction
	}
	distance := math.Max(float64(j-y), float64(y-j))
	distance = math.Max(1, distance)
	return int64(distance)
}
