package roman_converter

import (
	"fmt"
	"strings"
)

var romanNumerals = []struct {
	Value     int
	Symbol    string
	Occurance int
}{
	{1000, "M", 3},
	{900, "CM", 1},
	{500, "D", 1},
	{400, "CD", 1},
	{100, "C", 3},
	{90, "XC", 1},
	{50, "L", 1},
	{40, "XL", 1},
	{10, "X", 3},
	{9, "IX", 1},
	{5, "V", 1},
	{4, "IV", 1},
	{1, "I", 3},
}

func DecimalToRoman(decimal int) (string, error) {
	if decimal < 1 || decimal > 3999 {
		return "", fmt.Errorf("decimal value must be between 1 and 3999")
	}

	roman := ""

	for _, rn := range romanNumerals {
		for decimal >= rn.Value {
			roman += rn.Symbol
			decimal -= rn.Value
		}
	}

	return roman, nil
}

func RomanToDecimal(roman string) (int, error) {

	if roman == "" {
		return 0, fmt.Errorf("empty roman numeral")
	}

	decimal := 0
	for _, rn := range romanNumerals {
		sum := 0
		for sum < rn.Occurance {
			if strings.HasPrefix(roman, rn.Symbol) {
				decimal += rn.Value
				roman = roman[len(rn.Symbol):]
				sum++
			} else {
				break
			}
		}
	}

	if len(roman) > 0 {
		return 0, fmt.Errorf("invalid roman numeral: %s", roman)
	}

	return decimal, nil
}
