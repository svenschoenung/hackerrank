package main

import "fmt"
import "strings"

func main() {
	numbers := "0123456789"
	lower_case := "abcdefghijklmnopqrstuvwxyz"
	upper_case := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	special_characters := "!@#$%^&*()-+"

	var s_len int
	fmt.Scanf("%v\n", &s_len)

	var s string
	fmt.Scanf("%v\n", &s)

	missing_length := 6 - s_len

	missing_char_classes := 0
	if !strings.ContainsAny(s, numbers) {
		missing_char_classes++
	}
	if !strings.ContainsAny(s, lower_case) {
		missing_char_classes++
	}
	if !strings.ContainsAny(s, upper_case) {
		missing_char_classes++
	}
	if !strings.ContainsAny(s, special_characters) {
		missing_char_classes++
	}

	if missing_length > missing_char_classes {
		fmt.Println(missing_length)
	} else {
		fmt.Println(missing_char_classes)
	}
}
