package main

import (
	"fmt"
	"os"
	"strings"
)

func calculate(i int, j int, dx int, dy int, rows []string) bool {
	pattern := "MAS"
	for n := range 3 {
		newY := j + (dy * (n + 1))
		newX := i + (dx * (n + 1))

		if newX < 0 || newX >= len(rows) || newY < 0 || newY >= len(rows[0]) {
			return false
		}

		if rows[newX][newY] != pattern[n] {
			return false
		}
	}

	return true
}

func main() {
	data, _ := os.ReadFile("input.txt")

	input := strings.Split(string(data), "\n")

	var rows []string

	for _, v := range input {
		if strings.TrimSpace(v) != "" {

			rows = append(rows, strings.TrimSpace(v))
		}
	}

	count := 0

	for i, row := range rows {
		for j, char := range strings.Split(row, "") {
			if string(char) != "X" {
				continue
			}
			for _, dy := range [3]int{-1, 0, 1} {
				for _, dx := range [3]int{-1, 0, 1} {
					if calculate(i, j, dx, dy, rows) {
						count++
					}
				}
			}
		}
	}

	fmt.Println(count)

}
