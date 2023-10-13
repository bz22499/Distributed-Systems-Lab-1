package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func handleConnection(conn net.Conn) {
	reader := bufio.NewReader(conn)
	msg, _ := reader.ReadString('\n')
	fmt.Printf(msg)
}

func main() {
	stdin := bufio.NewReader(os.Stdin)
	connection, _ := net.Dial("tcp", "127.0.0.1:9999")
	for {
		fmt.Printf("Enter text:")
		msg, _ := stdin.ReadString('\n')
		fmt.Fprintln(connection, msg)
		handleConnection(connection)
	}
}
