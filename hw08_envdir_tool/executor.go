package main

import (
	"fmt"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	command := exec.Command(cmd[0], cmd[1:]...)
	command.Env = os.Environ()
	envSlice := []string{}
	for k, v := range env {
		if v.NeedRemove {
			// If NeedRemove is true, don't include this environment variable
			continue
		}
		envSlice = append(envSlice, fmt.Sprintf("%s=%s", k, v.Value))
	}
	command.Env = append(command.Env, envSlice...)

	err := command.Run()
	if e, ok := err.(*exec.ExitError); ok {
		return e.ExitCode()
	} else if err != nil {
		fmt.Println("Error executing command:", err)
		return 1
	}
	return 0
}
