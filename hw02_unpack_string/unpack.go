package hw02unpackstring

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var builder strings.Builder
	var result string
	count := utf8.RuneCountInString(str)

	if count == 0 {
		return "", nil
	}

	err := valid(str)
	if err != nil {
		return "", err
	}

	runes := []rune(str)
	for i, char := range runes {
		tt := []rune(str)[i]

		if unicode.IsDigit(tt) {
			i++
			pp := []rune(str)[i]
			if unicode.IsDigit(pp) {
				return "", ErrInvalidString
			}
			val, _ := strconv.Atoi(string(char))
			if val == 0 {
				res := zero(builder.String())
				builder.Reset()
				builder.WriteString(res)
			}

			for y := 0; y < val-1; y++ {
				ff := []rune(str)[i-2]
				builder.WriteString(string(ff))
				result += string(ff)
				continue
			}
		} else {
			builder.WriteString(string(char))
			result += string(char)
		}
	}
	fmt.Println(builder.String())
	return builder.String(), nil
}

func valid(str string) error {
	runes := []rune(str)
	for i, char := range runes {
		if i == 0 && unicode.IsDigit(char) {
			return ErrInvalidString
		}

		tt := []rune(str)[i]
		if unicode.IsDigit(tt) {
			i++
			pp := []rune(str)[i]
			if unicode.IsDigit(pp) {
				return ErrInvalidString
			}
		}
	}
	return nil
}

func zero(str string) string {
	var builder strings.Builder
	for i := 0; i < len(str)-1; i++ {
		builder.WriteString(string(str[i]))
	}
	return builder.String()
}
