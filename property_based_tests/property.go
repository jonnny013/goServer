package propertybasedtests

import "strings"

type RomanNumeral struct {
	Value  uint16
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func RomanNumeralConverter(num uint16) string {
	var romanNumeral strings.Builder
	for _, numerals := range allRomanNumerals {
		for num >= numerals.Value {
			romanNumeral.WriteString(numerals.Symbol)
			num -= numerals.Value
		}
	}
	return romanNumeral.String()
}

func ConvertToArabic(roman string) uint16 {
	var total uint16 = 0
	for _, characters := range allRomanNumerals {
		for strings.HasPrefix(roman, characters.Symbol) {
			total += characters.Value
			roman = strings.TrimPrefix(roman, characters.Symbol)
		}
	}
	return total
}
