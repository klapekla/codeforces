package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := ""
	fmt.Scanln(&input)
	summands := sliceOfStringsToSliceOfInts(strings.Split(input, "+"))
	sort.Ints(summands)
	new_sum := string(strings.Join(sliceOfIntsToSliceOfStrings(summands), "+"))
	fmt.Println(new_sum)
}

func sliceOfStringsToSliceOfInts(sliceOfStrings []string) []int {
	sliceOfInts := make([]int, len(sliceOfStrings))
	for i := range sliceOfStrings {
		sliceOfInts[i], _ = strconv.Atoi(sliceOfStrings[i])
	}
	return sliceOfInts
}

func sliceOfIntsToSliceOfStrings(sliceOfInts []int) []string {
	sliceOfStrings := make([]string, len(sliceOfInts))
	for i := range sliceOfInts {
		sliceOfStrings[i] = strconv.Itoa(sliceOfInts[i])
	}
	return sliceOfStrings
}
