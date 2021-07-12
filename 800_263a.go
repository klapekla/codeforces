package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	matrix := make([][]string, 5)
	scanner := bufio.NewScanner(os.Stdin)

	for i := range matrix {
		matrix[i] = make([]string, 5)
		scanner.Scan()
		input := scanner.Text()
		matrix[i] = strings.Split(input, " ")
	}

	// find position of 1
	var pos_row int
	var pos_col int
	for i, row := range matrix {
		for j, col := range row {
			if col == "1" {
				pos_row = int(i)
				pos_col = int(j)
			}
		}
	}

	// calculate moves
	var moves int
	if pos_row >= 2 {
		moves = pos_row - 2
	} else {
		moves = 2 - pos_row
	}
	if pos_col >= 2 {
		moves = moves + pos_col - 2
	} else {
		moves = moves + 2 - pos_col
	}

	fmt.Println(moves)
}
