package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var firstArray []int
	var secondArray []int
	var sum = 0

	b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var numbers = strings.Split(string(b), "\n")

	for _, v := range numbers {
		var segments = strings.Split(v, " ")
		i, _ := strconv.Atoi(strings.Trim(segments[0], " "))
		i2, _ := strconv.Atoi(strings.Trim(segments[len(segments)-1], " "))
		firstArray = append(firstArray, i)
		secondArray = append(secondArray, i2)
	}

	sort.Slice(firstArray, func(i, j int) bool { return firstArray[i] < firstArray[j] })
	sort.Slice(secondArray, func(i, j int) bool { return secondArray[i] < secondArray[j] })

	for i := 0; i < len(firstArray); i++ {
		var diff = firstArray[i] - secondArray[i]

		if diff > 0 {
			sum += diff
		} else {
			sum -= diff
		}
	}

	fmt.Println(sum)
}
