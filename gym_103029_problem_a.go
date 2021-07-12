package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var val []string

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	val = strings.Split(scanner.Text(), " ")
	type_of_nuts, _ := strconv.Atoi(val[0])
	nuts, _ := strconv.Atoi(val[1])

	scanner.Scan()
	val = strings.Split(scanner.Text(), " ")
	var nuts_in_box []int
	for _, v := range val {
		vi, _ := strconv.Atoi(v)
		nuts_in_box = append(nuts_in_box, vi)
	}

	scanner.Scan()
	_ = scanner.Text()

	scanner.Scan()
	val = strings.Split(scanner.Text(), " ")
	var nut_types_to_eat []int
	for _, v := range val {
		vi, _ := strconv.Atoi(v)
		nut_types_to_eat = append(nut_types_to_eat, vi)
	}

	fmt.Println("type_of_nuts:", type_of_nuts)
	fmt.Println("nuts:", nuts)
	fmt.Println("nuts_in_box:", nuts_in_box)
	fmt.Println("nut_types_to_eat:", nut_types_to_eat)
}
