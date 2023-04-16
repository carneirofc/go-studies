package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := strings.TrimSpace(input.Text())
		counts[line] += 1
	}
	err := input.Err()
	if err != nil {
		fmt.Printf("Error when reading input from stdin %v\n", err)
		os.Exit(-1)
	}
	for key, val := range counts {
		fmt.Printf("%d: %v\n", val, key)
	}
}
