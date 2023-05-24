package roman_converter

import (
	"testing"
)

type romanConversionTest struct {
	decimal int
	roman   string
	err     bool
}

var tests []romanConversionTest

func setup() {
	tests = []romanConversionTest{
		{1, "I", false},
		{4, "IV", false},
		{9, "IX", false},
		{49, "XLIX", false},
		{3999, "MMMCMXCIX", false},
		{5, "V", false},
		{10, "X", false},
		{50, "L", false},
		{100, "C", false},
		{500, "D", false},
		{1000, "M", false},
		{3, "III", false},
		{8, "VIII", false},
		{58, "LVIII", false},
		{88, "LXXXVIII", false},
		{1994, "MCMXCIV", false},
		{0, "", true},
		{4000, "MMMM", true},
	}
}

func TestDecimalToRoman(t *testing.T) {
	setup()

	for _, tc := range tests {
		result, err := DecimalToRoman(tc.decimal)

		if tc.err {
			if err == nil {
				t.Errorf("Expected error for decimal %d, but got nil", tc.decimal)
			}
		} else {
			if err != nil {
				t.Errorf("Error converting decimal %d to Roman numeral: %v", tc.decimal, err)
			} else if result != tc.roman {
				t.Errorf("DecimalToRoman(%d) = %s; want %s", tc.decimal, result, tc.roman)
			}
		}
	}
}

func TestRomanToDecimal(t *testing.T) {
	setup()

	for _, tc := range tests {
		result, err := RomanToDecimal(tc.roman)

		if tc.err {
			if err == nil {
				t.Errorf("Expected error for Roman numeral %s, but got nil", tc.roman)
			}
		} else {
			if err != nil {
				t.Errorf("Error converting Roman numeral %s to decimal: %v", tc.roman, err)
			} else if result != tc.decimal {
				t.Errorf("RomanToDecimal(%s) = %d; want %d", tc.roman, result, tc.decimal)
			}
		}
	}
}

func TestConversionRoundTrip(t *testing.T) {
	setup()

	for _, tt := range tests {
		result, err := RomanToDecimal(tt.roman)

		if tt.err {
			if err == nil {
				t.Errorf("Expected error for Roman numeral %s, but got nil", tt.roman)
			}
		} else {
			if err != nil {
				t.Errorf("Error converting Roman numeral %s to decimal: %v", tt.roman, err)
			} else if result != tt.decimal {
				t.Errorf("RomanToDecimal(%s) = %d; want %d", tt.roman, result, tt.decimal)
			}
		}
	}
}
