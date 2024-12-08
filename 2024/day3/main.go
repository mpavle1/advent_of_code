package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	data, _ := os.ReadFile("input2.txt")
	inputText := string(data)

	numberReg := regexp.MustCompile(`\d{1,3}`)
	mulReg := regexp.MustCompile(`mul\(\d{1,3}\,\d{1,3}\)`)

	matches := mulReg.FindAllString(inputText, -1)

	sum := 0
	for _, match := range matches {
		numbers := numberReg.FindAllString(match, -1)

		firstNum, _ := strconv.Atoi(numbers[0])
		secondNum, _ := strconv.Atoi(numbers[1])

		sum += firstNum * secondNum
	}

	fmt.Println(sum)
}
