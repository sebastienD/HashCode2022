package main

import (
	"bufio"
	"strconv"
	"strings"
)

func AtoI(v string) int {
	val, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}
	return val
}

func scanAndSplit(scanner *bufio.Scanner) []string {
	scanner.Scan()
	return strings.Split(scanner.Text(), " ")
}
