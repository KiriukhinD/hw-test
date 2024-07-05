package main

import "testing"

// Let's assume we have the following directory structure for environment variables:
// /path/to/env/dir/
// ├── var1
// ├── var2
// └── var3

func TestReadDir(t *testing.T) {
	expectedEnv := map[string]EnvValue{
		"var1": {"value1", false},
		"var2": {"value2", false},
		"var3": {"", true},
	}

	testEnvDir := "/path/to/env/dir"
	env, err := ReadDir(testEnvDir)
	if err != nil {
		t.Errorf("Error reading envdir: %v", err)
	}

	if len(env) != len(expectedEnv) {
		t.Errorf("Expected %d env variables, but got %d", len(expectedEnv), len(env))
	}

	for key, expectedValue := range expectedEnv {
		actualValue, exists := env[key]
		if !exists {
			t.Errorf("Expected env variable %s doesn't exist", key)
		} else {
			if actualValue.Value != expectedValue.Value {
				t.Errorf("Env variable %s has unexpected value. Expected: %s, Actual: %s", key, expectedValue.Value, actualValue.Value)
			}
			if actualValue.NeedRemove != expectedValue.NeedRemove {
				t.Errorf("NeedRemove flag for env variable %s is unexpected. Expected: %t, Actual: %t", key, expectedValue.NeedRemove, actualValue.NeedRemove)
			}
		}
	}
}
