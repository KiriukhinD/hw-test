package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestRunCmd(t *testing.T) {
	// Создаем временную директорию для тестирования
	tmpDir, err := ioutil.TempDir("", "envdir_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Создаем временные файлы окружения для тестирования
	env := map[string]EnvValue{
		"TEST_VAR1": {Value: "test value 1", NeedRemove: false},
		"TEST_VAR2": {Value: "test value 2", NeedRemove: false},
		"TEST_VAR3": {Value: "test value 3", NeedRemove: true},
	}

	for key, value := range env {
		filePath := filepath.Join(tmpDir, key)
		if value.NeedRemove {
			// Создаем пустой файл, если переменная должна быть удалена
			file, err := os.Create(filePath)
			if err != nil {
				t.Fatalf("Failed to create file: %v", err)
			}
			file.Close()
		} else {
			// Создаем файл с значением переменной
			if err := ioutil.WriteFile(filePath, []byte(value.Value), 0644); err != nil {
				t.Fatalf("Failed to write file: %v", err)
			}
		}
	}

	// Запускаем тестируемую функцию
	cmd := "echo"
	args := []string{"$TEST_VAR1", "$TEST_VAR2", "$TEST_VAR3"}
	err = execCmd(cmd, args, env)

	// Проверяем, что ошибки нет
	if err != nil {
		t.Errorf("execCmd failed: %v", err)
	}
}
