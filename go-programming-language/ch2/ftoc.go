package main

import "fmt"

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%0.4f°F = %0.4f°C\n", freezingF, fToC(freezingF))
	fmt.Printf("%0.4f°F = %0.4f°C\n", boilingF, fToC(boilingF))
}
