// Package luhn is Exercism.io exercise
package luhn

import (
	"strconv"
	"strings"
)

// Valid — Determine whether or not number string is valid per the Luhn formula.
func Valid(input string) bool {
	input = strings.ReplaceAll(input, " ", "")
	if len(input) <= 1 {
		return false
	}
	sum := 0
	isSecondDigit := len(input)%2 == 0
	for _, r := range input {
		v, err := strconv.Atoi(string(r))
		if err != nil {
			return false
		}
		if isSecondDigit {
			v *= 2
			if v > 9 {
				v -= 9
			}
		}
		sum += v
		isSecondDigit = !isSecondDigit
	}
	return sum%10 == 0
}

