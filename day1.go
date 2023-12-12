package main

import (
	"fmt"
	"strconv"
)

func day1() {
	lines := readStdin()

	answer := 0

	for _, line := range lines {
		first := 0
		last := 0
		foundFirst := false

		for i := 0; i < len(line); i++ {
			char := line[i]

			num, err := strconv.Atoi(string(char))

			if err == nil {
				if !foundFirst {
					first = num
					foundFirst = true
				}

				last = num
			}

		}

		s := fmt.Sprintf("%d%d", first, last)

		val, _ := strconv.Atoi(s)

		answer += val
	}

	fmt.Println(answer)
}
