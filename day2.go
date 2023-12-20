package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day2() {
	lines := readStdin()

	sum := 0

line_loop:
	for idx, line := range lines {
		input := strings.Split(line, ":")[1]

		sets := strings.Split(input, ";")

		for _, set := range sets {
			possible := evaluate_set(set)

			if !possible {
				continue line_loop
			}
		}

		sum += idx + 1
	}

	fmt.Println(sum)
}

func evaluate_set(set string) bool {
	cubes_all := strings.Split(set, ",")

	for _, cubes := range cubes_all {
		splitted_cube := strings.Split(cubes, " ")

		number, _ := strconv.Atoi(splitted_cube[1])
		color := splitted_cube[2]

		if color == "blue" && number > 14 {
			return false
		}

		if color == "red" && number > 12 {
			return false
		}

		if color == "green" && number > 13 {
			return false
		}
	}

	return true
}
