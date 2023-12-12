package main

import (
	"bufio"
	"fmt"
	"os"
)

func readStdin() []string {
	lines := []string{}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			break
		}

		lines = append(lines, text)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return lines
}
