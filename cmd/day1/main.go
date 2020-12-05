package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	entries, err := readIntegers("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Part 1
	const sum = 2020
	var n, entry1, entry2 int
Part1:
	for n, entry1 = range entries {
		for _, entry2 = range entries[n:] {
			if entry1+entry2 == sum {
				break Part1
			}
		}
	}

	fmt.Println(entry1)
	fmt.Println(entry2)
	fmt.Println(entry1 * entry2)

	// Part 2
	fmt.Println()
	var entry3 int
Part2:
	for n, entry1 = range entries {
		for _, entry2 = range entries[n:] {
			if entry1+entry2 < sum {
				for _, entry3 = range entries[n+1:] {
					if entry1+entry2+entry3 == sum {
						break Part2
					}
				}
			}
		}
	}

	fmt.Println(entry1)
	fmt.Println(entry2)
	fmt.Println(entry3)
	fmt.Println(entry1 * entry2 * entry3)

	os.Exit(0)
}

func readWords(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	return words, scanner.Err()
}

func readIntegers(filename string) ([]int, error) {
	words, err := readWords(filename)
	if err != nil {
		return nil, err
	}

	var ints []int
	for _, w := range words {
		i, err := strconv.Atoi(w)
		if err != nil {
			return nil, err
		}
		ints = append(ints, i)
	}

	return ints, nil
}
