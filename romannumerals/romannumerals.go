package romannumerals

import (
	"strings"
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

func Encode(num int) (string, bool) {
	if num < 1 {
		return "", false
	}
	u := num % 10
	t := (num%100 - u) / 10
	h := (num%1000 - t - u) / 100
	k := num / 1000

	sU := digitToString(u, []string{"I", "V", "X"})
	sT := digitToString(t, []string{"X", "L", "C"})
	sH := digitToString(h, []string{"C", "D", "M"})
	sK := digitToString(k, []string{"M", "", ""})

	return strings.Join([]string{sK, sH, sT, sU}, ""), true
}

func digitToString(digit int, chars []string) string {
	switch digit {
	case 1:
		return chars[0]
	case 2:
		return strings.Repeat(chars[0], 2)
	case 3:
		return strings.Repeat(chars[0], 3)
	case 4:
		return strings.Join([]string{chars[0], chars[1]}, "")
	case 5:
		return chars[1]
	case 6:
		return strings.Join([]string{chars[1], chars[0]}, "")
	case 7:
		return strings.Join([]string{chars[1], strings.Repeat(chars[0], 2)}, "")
	case 8:
		return strings.Join([]string{chars[1], strings.Repeat(chars[0], 3)}, "")
	case 9:
		return strings.Join([]string{chars[0], chars[2]}, "")
	default:
		return ""

	}
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
