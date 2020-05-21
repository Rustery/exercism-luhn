// Package luhn is Exercism.io exercise
package luhn

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// Valid — Determine whether or not number string is valid per the Luhn formula.
func Valid(input string) bool {
	cleanInput, err := clearInput(input)
	if err != nil {
		return false
	}
	reversedInput := reverse(cleanInput)
	sum := 0
	for k, val := range reversedInput {
		v, _ := strconv.Atoi(string(val))
		if k%2 == 0 {
			sum += int(v)
		} else {
			doubleNum := int(v) * 2
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

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func clearInput(input string) (string, error) {
	cleanInput := strings.Replace(input, " ", "", -1)
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		return "", err
	}
	if reg.MatchString(cleanInput) {
		return "", errors.New("Strings can contain only numbers and spaces")
	}
	if len(cleanInput) <= 1 {
		return "", errors.New("Strings of length 1 or less are not valid")
	}
	return cleanInput, nil
}
