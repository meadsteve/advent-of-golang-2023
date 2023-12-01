package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	//Read input file
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	var total = int64(0)

	for sc.Scan() {
		total = total + findFirstAndLastDigit(sc.Text())
	}
	fmt.Println(total)
}

func findFirstAndLastDigit(input string) int64 {
	first := findFirstDigit(input)
	last := findLastDigit(input)

	number, _ := strconv.ParseInt(string(*first)+string(*last), 10, 64)

	return number
}

func findFirstDigit(input string) *int32 {
	for _, ch := range input {
		if unicode.IsDigit(ch) {
			return &ch
		}
	}
	return nil
}

func findLastDigit(input string) *int32 {
	return findFirstDigit(reverse(input))
}

// reverse Stolen from stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go
func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
