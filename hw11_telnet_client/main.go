package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	// Определяем флаги
	timeout := flag.String("timeout", "10s", "Timeout for the connection")
	flag.Parse()

	// Получаем остаточные аргументы (хост и порт)
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: go-telnet --timeout=<duration> <host>:<port>")
		return
	}

	address := args[0]

	// Преобразуем timeout в значение типа time.Duration
	duration, err := time.ParseDuration(*timeout)
	if err != nil {
		fmt.Printf("Invalid timeout value: %v\n", err)
		return
	}

	// Открытие соединения с Telnet сервером
	client := NewTelnetClient(address, duration, os.Stdin, os.Stdout)
	err = client.Connect()
	if err != nil {
		fmt.Println("Error connecting to the Telnet server:", err)
		return
	}

	// Отправка данных на сервер
	_, err = fmt.Fprintln(os.Stdout, "Please enter your command:")
	if err != nil {
		return
	}
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
