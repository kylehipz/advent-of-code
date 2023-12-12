package main

import (
	"fmt"
	"strconv"
)

var numberMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"zero":  0,
}

func day2() {
	lines := readStdin()
	answer := 0

	for _, line := range lines {
		first := 0
		last := 0
		foundFirst := false

		for i := 0; i < len(line); i++ {
			char := line[i]

			// Check if current character can be converted to number
			num, err := strconv.Atoi(string(char))

			// If converted successfully
			if err == nil {
				if !foundFirst {
					first = num
					foundFirst = true
				}

				last = num
			} else {
				// Slice the line from the current character up to 5 characters
				// and check if it is a digit word
				for j := 1; j <= 5; j++ {
					if i+j > len(line) {
						continue
					}

					potentialDigitWord := line[i : i+j]
					potentialDigit, ok := numberMap[potentialDigitWord]

					// word is a digit
					if ok {
						if !foundFirst {
							first = potentialDigit
							foundFirst = true
						}

						last = potentialDigit
					}
				}
			}
		}

		s := fmt.Sprintf("%d%d", first, last)
		fmt.Println(first, last)

		val, _ := strconv.Atoi(s)

		answer += val
	}

	fmt.Println(answer)
}
