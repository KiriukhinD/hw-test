package main

import (
	"io/ioutil"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dirPath string) (map[string]EnvValue, error) {
	env := make(map[string]EnvValue)

	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		key := file.Name()
		value, err := ioutil.ReadFile(dirPath + "\\" + key)
		if err != nil {
			return nil, err
		}
		env[key] = EnvValue{Value: strings.TrimSpace(string(value)), NeedRemove: false}
	}

	return env, nil
}
