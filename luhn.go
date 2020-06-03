// Package luhn is Exercism.io exercise
package luhn

import (
	"strconv"
	"strings"
)

// Valid — Determine whether or not number string is valid per the Luhn formula.
func Valid(input string) bool {
	cleanInput := strings.ReplaceAll(input, " ", "")
	if len(cleanInput) <= 1 {
		return false
	}
	sum := 0
	for i, isSecondDigit := len(cleanInput), true; i > 0; i, isSecondDigit = i-1 , !isSecondDigit {
		v, err := strconv.Atoi(string(cleanInput[i-1]))
		if err != nil {
			return false
		}
		if isSecondDigit {
			sum += v
		} else {
			doubleNum := v * 2
			if doubleNum > 9 {
				sum += doubleNum - 9
			} else {
				sum += doubleNum
			}
		}
	}
	if sum%10 == 0 {
		return true
	}
	return false
}
