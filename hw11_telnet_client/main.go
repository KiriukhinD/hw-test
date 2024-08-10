package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// Проверка правильности количества аргументов
	if len(os.Args) < 2 {
		fmt.Println("Usage: go-telnet <host>:<port>")
		return
	}

	address := os.Args[1]
	timeout := 10 * time.Second // Установка тайм-аута в 10 секунд

	// Открытие соединения с Telnet сервером
	client := NewTelnetClient(address, timeout, os.Stdin, os.Stdout)
	err := client.Connect()
	if err != nil {
		fmt.Println("Error connecting to the Telnet server:", err)
		return
	}
	defer client.Close()

	// Отправка данных на сервер
	fmt.Fprintln(os.Stdout, "Please enter your command:")
	err = client.Send()
	if err != nil {
		fmt.Println("Error sending data:", err)
		return
	}

	// Получение ответа от сервера
	fmt.Println("Response from the server:")
	err = client.Receive()
	if err != nil {
		fmt.Println("Error receiving data:", err)
		return
	}
}
