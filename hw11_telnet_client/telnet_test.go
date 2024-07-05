package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"net"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTelnetClient(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		l, err := net.Listen("tcp", "127.0.0.1:")
		require.NoError(t, err)
		defer func() { require.NoError(t, l.Close()) }()

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()

			in := &bytes.Buffer{}
			out := &bytes.Buffer{}

			timeout, err := time.ParseDuration("10s")
			require.NoError(t, err)

			client := NewTelnetClient(l.Addr().String(), timeout, io.NopCloser(in), out)
			require.NoError(t, client.Connect())
			defer func() { require.NoError(t, client.Close()) }()

			in.WriteString("hello\n")
			err = client.Send()
			require.NoError(t, err)

			err = client.Receive()
			require.NoError(t, err)
			require.Equal(t, "world\n", out.String())
		}()

		go func() {
			defer wg.Done()

			conn, err := l.Accept()
			require.NoError(t, err)
			require.NotNil(t, conn)
			defer func() { require.NoError(t, conn.Close()) }()

			request := make([]byte, 1024)
			n, err := conn.Read(request)
			require.NoError(t, err)
			require.Equal(t, "hello\n", string(request)[:n])

			n, err = conn.Write([]byte("world\n"))
			require.NoError(t, err)
			require.NotEqual(t, 0, n)
		}()

		wg.Wait()
	})
}
func TestConnect(t *testing.T) {
	address := "example.com:23"
	timeout := 10 * time.Second
	in := strings.NewReader("test input")
	out := &bytes.Buffer{}
	client := NewTelnetClient(address, timeout, ioutil.NopCloser(in), out)

	err := client.Connect()
	if err != nil {
		t.Errorf("Error connecting: %v", err)
	}

	if client.Connect() == nil {
		t.Error("Connection was not established")
	}

	client.Close()
}

func TestSendReceive(t *testing.T) {
	address := "example.com:23"
	timeout := 10 * time.Second
	in := strings.NewReader("test input")
	out := &bytes.Buffer{}
	client := NewTelnetClient(address, timeout, ioutil.NopCloser(in), out)

	conn, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("Error creating mock server: %v", err)
	}
	defer conn.Close()

	go func() {
		for {
			client, err := conn.Accept()
			if err != nil {
				return
			}
			defer client.Close()
			io.Copy(client, client)
		}
	}()

	err = client.Connect()
	if err != nil {
		t.Fatalf("Error connecting: %v", err)
	}

	err = client.Send()
	if err != nil {
		t.Fatalf("Error sending: %v", err)
	}

	err = client.Receive()
	if err != nil {
		t.Fatalf("Error receiving: %v", err)
	}
}
