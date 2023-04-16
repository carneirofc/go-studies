package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLines(file *os.File, counts map[string]int) error {
	input := bufio.NewScanner(file)
	for input.Scan() {
		line := input.Text()
		counts[line] += 1
	}
	err := input.Err()
	if err != nil {
		return err
	}

	return nil
}

func printCounts(counts map[string]int) {
	for key, val := range counts {
		fmt.Printf("%d: %v\n", val, key)
	}
}

func main() {
	files := os.Args[1:]

	counts := make(map[string]int)

	if len(files) == 0 {
		err := countLines(os.Stdin, counts)
		if err != nil {
			fmt.Printf("Error when reading input from stdin %v\n", err)
			os.Exit(-1)
		}
	}

	for _, file := range files {
		fd, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file '%s' '%v'\n", file, err)
			os.Exit(-1)
		}
		err = countLines(fd, counts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when reading input from stdin %v\n", err)
			os.Exit(-1)
		}
	}
	printCounts(counts)
}
