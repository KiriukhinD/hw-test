package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go-envdir /path/to/env/dir command arg1 arg2 ...")
		os.Exit(1)
	}

	envDir := os.Args[1]
	cmd := os.Args[2]
	cmdArgs := os.Args[2:]
	env, err := ReadDir(envDir)
	if err != nil {
		fmt.Println("Error reading envdir:", err)
		os.Exit(1)
	}

	err = execCmd(cmd, cmdArgs, env)
	if err != nil {
		fmt.Println("Error executing command:", err)
		os.Exit(1)
	}
}

func processEnvFile(filePath string) (string, bool, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", false, err
	}

	contentStr := string(content)
	if len(contentStr) == 0 {
		return "", true, nil
	}

	value := strings.TrimRight(contentStr, " \t")
	value = strings.Replace(value, "\x00", "\n", -1)
	return value, false, nil
}

func execCmd(cmd string, args []string, env map[string]EnvValue) error {
	command := exec.Command(cmd, args...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Stdin = os.Stdin

	cmdEnv := os.Environ()
	for k, v := range env {
		if v.NeedRemove {
			cmdEnv = removeEnvVar(cmdEnv, k)
		} else {
			cmdEnv = append(cmdEnv, fmt.Sprintf("%s=%s", k, v.Value))
		}
	}
	command.Env = cmdEnv

	return command.Run()
}

func removeEnvVar(env []string, key string) []string {
	var newEnv []string
	for _, e := range env {
		if !strings.HasPrefix(e, key+"=") {
			newEnv = append(newEnv, e)
		}
	}
	return newEnv
}
