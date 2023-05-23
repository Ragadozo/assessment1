package main

import (
	"testing"
)

type romanConversionTest struct {
	decimal int
	roman   string
}

var tests []romanConversionTest

func setup() {
	tests = []romanConversionTest{
		{1, "I"},
		{4, "IV"},
		{9, "IX"},
		{49, "XLIX"},
		{3999, "MMMCMXCIX"},
	}
}

func TestDecimalToRoman(t *testing.T) {
	setup()

	for _, tt := range tests {
		result := DecimalToRoman(tt.decimal)
		if result != tt.roman {
			t.Errorf("DecimalToRoman(%d) = %s; want %s", tt.decimal, result, tt.roman)
		}
	}
}

func TestRomanToDecimal(t *testing.T) {
	setup()

	for _, tt := range tests {
		result := RomanToDecimal(tt.roman)
		if result != tt.decimal {
			t.Errorf("RomanToDecimal(%s) = %d; want %d", tt.roman, result, tt.decimal)
		}
	}
}

func TestConversionRoundTrip(t *testing.T) {
	setup()

	for _, tt := range tests {
		roman := DecimalToRoman(tt.decimal)
		result := RomanToDecimal(roman)
		if result != tt.decimal {
			t.Errorf("Conversion round trip failed for decimal %d. Got Roman: %s, Decimal: %d", tt.decimal, roman, result)
		}
	}
}
