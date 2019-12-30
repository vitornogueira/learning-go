package iteration

import "strings"

// Repeat a new string consisting of count copies of the string s
func Repeat(character string, repeatCount int) string {
	return strings.Repeat(character, repeatCount)
}
