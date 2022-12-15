package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type InputFunc func(string)

func ReadFile(fileName string, fun InputFunc) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		text := strings.TrimSpace(line)
		fun(text)
	}
	return nil
}

func PrettyPrintNum(num int64) (result string) {
	strNum := fmt.Sprintf("%v", num)
	length := len(strNum)
	for i := 1; i <= length; i++ {
		result = string(strNum[length-i]) + result
		if i > 0 && i%3 == 0 {
			result = "," + result
		}
	}
	return result
}
