package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	data, _ := os.ReadFile("input2.txt")
	parsedText := string(data)

	numberReg := regexp.MustCompile(`\d{1,3}`)
	reg := regexp.MustCompile(`mul\(\d{1,3}\,\d{1,3}\)`)

	matches := reg.FindAllString(parsedText, -1)

	sum := 0
	for i := 0; i < len(matches); i++ {
		numberMatches := numberReg.FindAllString(matches[i], -1)

		firstNum, _ := strconv.Atoi(numberMatches[0])
		secondNum, _ := strconv.Atoi(numberMatches[1])

		sum += firstNum * secondNum
	}

	fmt.Println(sum)
}
