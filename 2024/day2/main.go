package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")

	rows := strings.Split(string(data), "\n")
	var safeRows = 0

	for i := 0; i < len(rows)-1; i++ {
		sums := strings.Split(rows[i], " ")

		var hasInc = false
		var hasDec = false
		var hasBadDiff = false

		for j := 0; j < len(sums)-1; j++ {
			num1, _ := strconv.Atoi(sums[j])
			num2, _ := strconv.Atoi(sums[j+1])

			diff := num1 - num2

			// fmt.Println(diff, num1, num2)

			if diff > 0 {
				hasDec = true
			}
			if diff < 0 {
				hasInc = true
			}

			if diff < -3 || diff > 3 || diff == 0 {
				hasBadDiff = true
			}
		}

		// fmt.Println(sums, hasInc, hasDec, hasBadDiff)

		if !hasBadDiff && hasInc != hasDec {
			fmt.Println(sums)
			safeRows++
		}
	}
	fmt.Println(safeRows)
}
