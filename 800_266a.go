package main

import "fmt"

func main() {
	stones := 0
	colors := ""

	fmt.Scan(&stones)
	fmt.Scan(&colors)

	colorsInBytes := []byte(colors)
	answer := 0
	for i := 0; i < len(colorsInBytes)-1; i++ {
		if colorsInBytes[i] == colorsInBytes[i+1] {
			answer++
		}
	}

	fmt.Println(answer)
}
