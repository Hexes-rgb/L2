package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "timeout for connection")
	flag.Parse()

	if flag.NArg() < 2 {
		fmt.Println("Usage: task [--timeout=10s] host port")
		os.Exit(1)
	}

	host := flag.Arg(0)
	port := flag.Arg(1)
	address := net.JoinHostPort(host, port)

	conn, err := net.DialTimeout("tcp", address, *timeout)
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Connected to", address)

	done := make(chan struct{})

	go func() {
		io.Copy(os.Stdout, conn)
		done <- struct{}{}
	}()

	go func() {
		io.Copy(conn, os.Stdin)
		done <- struct{}{}
	}()

	<-done
	conn.Close()
	<-done
}
