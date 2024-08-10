package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	t.Run("check name", checkNameTests)
	t.Run("check handle value", checkHandleValueTests)
	t.Run("read value from file", readValueFromFileTests)
	t.Run("read dir", readDirTests)
}

func checkNameTests(t *testing.T) {
	t.Helper()
	require.True(t, checkName("asd"))
	require.False(t, checkName("VA=R"))
	require.False(t, checkName("VAЯR"))
	require.False(t, checkName(" VAR"))
}

func checkHandleValueTests(t *testing.T) {
	t.Helper()
	require.True(t, handleValue([]byte("")).NeedRemove)
	require.False(t, handleValue([]byte("hello")).NeedRemove)
	require.Equal(t, &EnvValue{Value: "hel\nlo", NeedRemove: false}, handleValue([]byte("hel\x00lo")))
	require.Equal(t, &EnvValue{Value: "  hello", NeedRemove: false}, handleValue([]byte("  hello  \t  ")))
}

func readValueFromFileTests(t *testing.T) {
	t.Helper()

	_, err := readValueFromFile("/dev", "nul=l")
	require.ErrorIs(t, err, ErrWrongVarName)

	_, err = readValueFromFile("/etc", "nonexistentfile123")
	require.ErrorIs(t, err, ErrUnableToOpenFile)

	tests := []struct {
		key      string
		expected EnvValue
	}{
		{"BAR", EnvValue{Value: "bar", NeedRemove: false}},
		{"EMPTY", EnvValue{Value: "", NeedRemove: true}},
		{"FOO", EnvValue{Value: "   foo\nwith new line", NeedRemove: false}},
		{"HELLO", EnvValue{Value: "\"hello\"", NeedRemove: false}},
		{"UNSET", EnvValue{Value: "", NeedRemove: true}},
	}

	for _, tt := range tests {
		result, err := readValueFromFile("testdata/env", tt.key)
		require.NoError(t, err)
		require.Equal(t, tt.expected, *result)
	}
}

func readDirTests(t *testing.T) {
	t.Helper()

	_, err := ReadDir("/notfounddir")
	require.ErrorIs(t, err, ErrUnableToReadDir)

	_, err = ReadDir("testdata/env/UNSET")
	require.ErrorIs(t, err, ErrNotADir)

	_, err = ReadDir("testdata/")
	require.ErrorIs(t, err, ErrNotAFile)

	_, err = ReadDir("/root")
	require.ErrorIs(t, err, ErrUnableToReadDir)

	result, _ := ReadDir("testdata/env")
	require.Equal(t, 5, len(result))
}
