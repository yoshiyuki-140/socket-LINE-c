package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	ip, port := "192.168.3.7", "8000"
	conn, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Start a goroutine to handle incoming messages
	go func() {
		reader := bufio.NewReader(conn)
		for {
			message, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading from server:", err)
				return
			}
			fmt.Print(message)
		}
	}()

	// Read messages from stdin and send to server
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter message: ")
		message, _ := reader.ReadString('\n')
		conn.Write([]byte(message))
	}
}
