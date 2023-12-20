package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day2Part2() {
	lines := readStdin()

	sum := 0

	for _, line := range lines {
		input := strings.Split(line, ":")[1]

		sets := strings.Split(input, ";")

		max_blue := 0
		max_red := 0
		max_green := 0
		for _, set := range sets {
			evaluateSet(set, &max_blue, &max_red, &max_green)
		}

		power := max_blue * max_red * max_green
		sum += power
	}

	fmt.Println(sum)
}

func evaluateSet(set string, max_blue *int, max_red *int, max_green *int) {
	cubes_all := strings.Split(set, ",")

	for _, cubes := range cubes_all {
		splitted_cube := strings.Split(cubes, " ")

		number, _ := strconv.Atoi(splitted_cube[1])
		color := splitted_cube[2]

		if color == "blue" && number > *max_blue {
			*max_blue = number
		}

		if color == "red" && number > *max_red {
			*max_red = number
		}

		if color == "green" && number > *max_green {
			*max_green = number
		}
	}
}
