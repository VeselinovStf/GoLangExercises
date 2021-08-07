// Package romannumerals implements functions over Roman Numbers
package romannumerals

import "errors"

var thousents = []string{
	"M",
	"MM",
	"MMM",
}

var hundrets = []string{
	"C",
	"CC",
	"CCC",
	"CD",
	"D",
	"DC",
	"DCC",
	"DCCC",
	"CM",
	"M",
}

var tens = []string{
	"X",
	"XX",
	"XXX",
	"XL",
	"L",
	"LX",
	"LXX",
	"LXXX",
	"XC",
	"C",
}

var nums = []string{
	"I",
	"II",
	"III",
	"IV",
	"V",
	"VI",
	"VII",
	"VIII",
	"IX",
	"X",
}

// ToRomanNumeral  converts from normal (arabic) numbers to Roman Numerals
func ToRomanNumeral(arabic int) (rom string, err error) {
	if arabic <= 0 || arabic > 3000 {
		return "", errors.New("try with possitive number please")
	}

	EXIT:
	for arabic > 0 {
		switch {
		case arabic >= 1000:
			r := extract(arabic, 1000)
			rom += thousents[r-1]
			arabic = arabic % 1000
		case arabic >= 100:
			r := extractMod(arabic, 100, 10)
			rom += hundrets[r-1]
			arabic = arabic % 100
		case arabic >= 10:
			r := extractMod(arabic, 10, 10)
			rom += tens[r-1]
			arabic = arabic % 10
		case arabic > 0:
			r := arabic % 10
			rom += nums[r-1]
			break EXIT
		}
	}

	return rom, nil
}

func extract(n, exp int) int {
	return (n / exp)
}

func extractMod(n, exp, mod int) int {
	return (n / exp) % mod
}
