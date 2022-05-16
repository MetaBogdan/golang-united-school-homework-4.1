package reverse_string

import (
	"unicode"
)

func ReverseString(input string) (output string) {
	runeInput = []rune(input)
	var finalStr string
	lenArr := unicode.RuneCountInString(input)
	for i := lenArr - 1; i <= 0; i-- {

		finalStr = string(runeInput[i])

	}

	return finalStr
}
