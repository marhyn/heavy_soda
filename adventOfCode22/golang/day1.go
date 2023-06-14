package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	readFile, err := os.Open("../resources/day1.txt")

	check(err)

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var elfCalories []int
	currentElf := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if line == "" {
			elfCalories = append(elfCalories, currentElf)
			currentElf = 0
		} else {
			calories, err := strconv.Atoi(line)
			check(err)

			currentElf += calories
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elfCalories)))
	fmt.Println("1st: ", elfCalories[0])
	fmt.Println("2st: ", elfCalories[0]+elfCalories[1]+elfCalories[2])

	readFile.Close()
}
