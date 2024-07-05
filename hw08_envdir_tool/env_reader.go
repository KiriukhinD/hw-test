package main

import (
	"io/ioutil"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	env := make(Environment)

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		filePath := dir + "/" + entry.Name()
		value, needRemove, err := processEnvFile(filePath)
		if err != nil {
			return nil, err
		}
		env[entry.Name()] = EnvValue{Value: value, NeedRemove: needRemove}
	}
	return env, nil
}
