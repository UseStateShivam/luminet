// Create a TCP client that connects to localhost:8000.
// Send a test message and read the response.
// Keep the connection open for continuous communication.

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		fmt.Printf("error connecting to 8000: %s", err)
		os.Exit(1)
	}
	defer conn.Close()
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Type your message to send:")

	for {
		str, _ := reader.ReadString('\n')
		conn.Write([]byte(str))

		reply, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Printf("server replied with: %s", reply)
	}
}