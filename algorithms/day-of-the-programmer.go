package main

import "fmt"

func daysInFeb(year int) int {
	if year == 1918 {
		if year%4 == 0 {
			return 16
		} else {
			return 15
		}
	} else if year < 1918 {
		if year%4 == 0 {
			return 29
		} else {
			return 28
		}
	} else {
		if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
			return 29
		} else {
			return 28
		}
	}
}

func main() {
	var year int
	fmt.Scanf("%v\n", &year)

	var daysInMonth = [12]int{
		31,
		daysInFeb(year),
		31,
		30,
		31,
		30,
		31,
		31,
		30,
		31,
		30,
		31,
	}

	month := 0
	days := 0
	for ; month <= 12; month++ {
		if days+daysInMonth[month] > 256 {
			break
		}
		days += daysInMonth[month]
	}

	day := 256 - days
	fmt.Printf("%02d.%02d.%d\n", day, month+1, year)
}
