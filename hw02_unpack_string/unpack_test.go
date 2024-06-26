package hw02unpackstring

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		{input: "а1б2в3", expected: "аббввв"},
		{input: "a0", expected: ""},
		{input: "a1", expected: "a"},
		{input: "я9", expected: "яяяяяяяяя"},
		{input: "สวัสดี", expected: "สวัสดี"},
		// uncomment if task with asterisk completed
		// {input: `qwe\4\5`, expected: `qwe45`},
		// {input: `qwe\45`, expected: `qwe44444`},
		// {input: `qwe\\5`, expected: `qwe\\\\\`},
		// {input: `qwe\\\3`, expected: `qwe\3`},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{"3abc", "45", "aaa10b"}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}

func TestZero(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"hello", "hell"},
		{"world", "worl"},
		{"", ""},
		{"12345", "1234"},
		{"abcdef", "abcde"},
		{"a", ""},
	}

	for _, tc := range testCases {
		result := zero(tc.input)
		if result != tc.expected {
			t.Errorf("For input %q, expected %q, but got %q", tc.input, tc.expected, result)
		}
	}
}
