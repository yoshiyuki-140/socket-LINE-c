package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
)

var (
    clients   = make(map[net.Conn]bool)
    messages  []string
    mutex     sync.Mutex
)

func handleConnection(conn net.Conn) {
    defer conn.Close()
    fmt.Println("Client connected:", conn.RemoteAddr())

    // Send existing messages to the new client
    for _, msg := range messages {
        conn.Write([]byte(msg + "\n"))
    }

    // Add client to the list
    mutex.Lock()
    clients[conn] = true
    mutex.Unlock()

    reader := bufio.NewReader(conn)
    for {
        message, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("Client disconnected:", conn.RemoteAddr())
            break
        }

        fmt.Print("Message received:", message)

        // Save message and broadcast to all clients
        mutex.Lock()
        messages = append(messages, message)
        for client := range clients {
            client.Write([]byte(message))
        }
        mutex.Unlock()
    }

    // Remove client from the list
    mutex.Lock()
    delete(clients, conn)
    mutex.Unlock()
}

func main() {
    listener, err := net.Listen("tcp", ":8000")
    if err != nil {
        fmt.Println("Error starting server:", err)
        os.Exit(1)
    }
    defer listener.Close()
    fmt.Println("Server listening on port 8000")

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting connection:", err)
            continue
        }
        go handleConnection(conn)
    }
}
