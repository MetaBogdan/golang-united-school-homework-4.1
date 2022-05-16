package reverse_string

import (
	"unicode/utf8"
)

func ReverseString(input string) (output string) {
	runeInput := []rune(input)
	var finalStr string
	lenArr := utf8.RuneCountInString(input)
	for i := lenArr - 1; i >= 0; i-- {
		finalStr += string(runeInput[i])
	}
	return finalStr
}
