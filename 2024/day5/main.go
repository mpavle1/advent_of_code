package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")

	segments := strings.Split(string(data), "\n\n")

	pairs := segments[0]
	inputs := strings.Split(segments[1], "\n")
	// Mozda je potrebno zbog toga sto slice iznad pokupi i prazan red na kraju
	inputs = inputs[:len(inputs)-1]

	sum := 0

	for _, row := range inputs {
		values := strings.Split(row, ",")

		isRowCorrect := true

		for i := 0; i < len(values)-1; i++ {
			for j := i + 1; j < len(values); j++ {
				needle := string(values[i]) + `\|` + string(values[j])
				r := regexp.MustCompile(needle)
				matches := r.FindAllStringSubmatch(pairs, -1)

				if len(matches) == 0 {
					isRowCorrect = false
				}
			}
		}

		if isRowCorrect {
			index := math.Ceil(float64(len(values) / 2))
			val, _ := strconv.Atoi(values[int(index)])
			sum += val
		}
	}

	fmt.Println(sum)
}
