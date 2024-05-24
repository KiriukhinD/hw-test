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

	runes := []rune(str)

	err := valid(str)
	if err != nil {
		return "", err
	}
	for i := 0; i < count; i++ {
		if unicode.IsDigit(runes[i]) {
			val, _ := strconv.Atoi(string(runes[i]))

			if val == 0 {
				res := zero(builder.String())
				builder.Reset()
				builder.WriteString(res)
			}

			for y := 0; y < val-1; y++ {
				builder.WriteString(string(runes[i-1]))
				result += string(runes[i-1])
				continue
			}
		} else {
			builder.WriteString(string(runes[i]))
			result += string(runes[i])
		}
	}

	fmt.Println(builder.String())
	return builder.String(), nil
}

func zero(str string) string {
	var builder strings.Builder
	for i := 0; i < len(str)-1; i++ {
		builder.WriteString(string(str[i]))
	}
	return builder.String()
}

func valid(str string) error {
	count := utf8.RuneCountInString(str)
	runes := []rune(str)

	for i := 0; i < count; i++ {
		if i == 0 && unicode.IsDigit(runes[i]) {
			return ErrInvalidString
		}

		if unicode.IsDigit(runes[i]) && unicode.IsDigit(runes[i-1]) {
			return ErrInvalidString
		}
	}
	return nil
}
