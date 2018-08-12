package main

import "fmt"

func main() {
	var slen int
	fmt.Scanf("%v\n", &slen)

	var s string
	fmt.Scanf("%v\n", &s)

	var offset int
	fmt.Scanf("%v\n", &offset)

	enc := ""
	for _, char := range s {
		code := int(char)
		if code >= 65 && code <= 90 {
			enc += string(rune((((code - 65) + offset) % 26) + 65))
		} else if code >= 97 && code <= 122 {
			enc += string(rune((((code - 97) + offset) % 26) + 97))
		} else {
			enc += string(char)
		}
	}
	fmt.Println(enc)
}
