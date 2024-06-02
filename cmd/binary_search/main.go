package main

import (
	binary_search "algorithms_go/pkg/fundamentals"
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("No arguments were provided.")
	}

	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var whitelist []int

	for scanner.Scan() {
		num, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if err != nil {
			fmt.Println("Error parsing integer:", err)
			return
		}
		whitelist = append(whitelist, num)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	slices.Sort(whitelist)

	stdinScanner := bufio.NewScanner(os.Stdin)

	for stdinScanner.Scan() {
		num, err := strconv.Atoi(strings.TrimSpace(stdinScanner.Text()))
		if err != nil {
			fmt.Println("Error parsing integer:", err)
			return
		}

		if binary_search.Rank(num, whitelist) == -1 {
			fmt.Println(num)
		}
	}

	if err := stdinScanner.Err(); err != nil {
		fmt.Println("Error reading from stdin:", err)
		return
	}
}
