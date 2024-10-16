package romannumerals

import (
	"fmt"
	"strconv"
)

var romanIntValues = map[string]int{
	"M": 1000,
	"D": 500,
	"C": 100,
	"L": 50,
	"X": 10,
	"V": 5,
	"I": 1,
}

var integerRomValues = map[string]string{
	"1": "I",
	"2": "500",
}

func Encode(num int) (rom string, ok bool) {
	if num < 1 {
		return "", false
	}
	numString := strconv.Itoa(num)

	// romString := ""
	switch numLen := len(numString); {
	case numLen > 3: // thousands
		fmt.Println(num, "thousands")
	case numLen > 2: // hundreds
		fmt.Println(num, "hundreds")
	case numLen > 1: // tens
		fmt.Println(num, "tens")
	default: // units
		fmt.Println(num, "units")

	}
	return "", false
}

func Decode(rom string) (int, bool) {
	num := 0
	ok := false

	for i, s := range rom {
		var intValue int
		romNumeral := string(s)

		// read integer value of roman letter from map
		intValue, ok = romanIntValues[romNumeral]
		if !ok {
			return 0, ok
		}

		// check for subtractive preceding letter
		if i > 0 {
			switch romNumeral {
			case "V", "X":
				if string(rom[i-1]) == "I" {
					intValue -= 2
				}
			case "L", "C":
				if string(rom[i-1]) == "X" {
					intValue -= 20
				}
			case "D", "M":
				if string(rom[i-1]) == "C" {
					intValue -= 200
				}
			}
		}

		// accumulate value
		num += intValue
	}

	return num, ok
}
