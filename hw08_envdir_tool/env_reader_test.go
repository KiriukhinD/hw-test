package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	t.Run("read value from file", func(t *testing.T) {
		_, err := readValueFromFile("/dev", "nul=l")
		require.ErrorIs(t, err, ErrWrongVarName)

		_, err = readValueFromFile("/etc", "nonexistentfile123")
		require.ErrorIs(t, err, ErrUnableToOpenFile)

		result, err := readValueFromFile("testdata/env", "BAR")
		require.NoError(t, err)
		require.Equal(t, EnvValue{
			Value:      "bar",
			NeedRemove: false,
		}, *result)

		result, err = readValueFromFile("testdata/env", "EMPTY")
		require.NoError(t, err)
		require.Equal(t, EnvValue{
			Value:      "",
			NeedRemove: true,
		}, *result)

		result, err = readValueFromFile("testdata/env", "FOO")
		require.NoError(t, err)
		require.Equal(t, EnvValue{
			Value:      "   foo\nwith new line",
			NeedRemove: false,
		}, *result)

		result, err = readValueFromFile("testdata/env", "HELLO")
		require.NoError(t, err)
		require.Equal(t, EnvValue{
			Value:      "\"hello\"",
			NeedRemove: false,
		}, *result)

		result, err = readValueFromFile("testdata/env", "UNSET")
		require.NoError(t, err)
		require.Equal(t, EnvValue{
			Value:      "",
			NeedRemove: true,
		}, *result)
	})
}
