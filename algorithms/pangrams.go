package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

func pangrams(s string) string {
	usedChars := map[rune]bool{}
	used := 0

	for _, char := range s {
		if char != ' ' {
			if !usedChars[unicode.ToLower(char)] {
				usedChars[unicode.ToLower(char)] = true
				used++
			}
		}
	}
	if used == 26 {
		return "pangram"
	} else {
		return "not pangram"
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	s := readLine(reader)

	result := pangrams(s)

	fmt.Println(result)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
