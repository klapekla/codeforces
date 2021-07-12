package main

import (
	"fmt"
	"strings"
)

func main() {
	input := ""
	fmt.Scanln(&input)
	output := strings.Title(input)
	fmt.Println(output)
}
