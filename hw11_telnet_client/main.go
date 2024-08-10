package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// ParseInput разбирает входную строку и возвращает команду и её аргументы
func ParseInput(input string) (string, []string) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return "", nil
	}
	command := parts[0]
	args := parts[1:]
	return command, args
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go-telnet <host>:<port>")
		return
	}

	address := os.Args[1]
	timeout := 10 * time.Second // Установка тайм-аута в 10 секунд

	client := NewTelnetClient(address, timeout, os.Stdin, os.Stdout)

	err := client.Connect()
	if err != nil {
		fmt.Println("Error connecting to the Telnet server:", err)
		return
	}

	// Чтение пользовательского ввода
	reader := bufio.NewReader(os.Stdin)
	_, err = fmt.Fprintln(os.Stdout, "Введите команду для сервера (или 'exit' для выхода):")
	if err != nil {
		return
	}

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')

		// Удаляем символы новой строки
		input = strings.TrimSpace(input)

		// Проверка команды выхода
		if input == "exit" {
			break
		}

		// Разбор команды и аргументов
		command, args := ParseInput(input)

		// Логируем отправляемую команду
		if len(args) > 0 {
			fmt.Printf("Отправка команды: %s, аргументы: %v\n", command, args)
		} else {
			fmt.Printf("Отправка команды: %s\n", command)
		}

		// Отправка команды на сервер
		err = client.Send()
		if err != nil {
			fmt.Println("Ошибка при отправке данных:", err)
			return
		}

		// Получение ответа от сервера
		fmt.Println("Ответ от сервера:")
		err = client.Receive()
		if err != nil {
			fmt.Println("Ошибка при получении данных:", err)
			return
		}
	}
}
