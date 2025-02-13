// Using Go’s net package to listen on port 8000.
// Accept incoming connections and read data using a loop.
// Print received messages and send a response back to the client.

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Printf("error in listening on 8000: %s", err)
		os.Exit(1)
	}

	// 	defer is used to schedule a function call to be executed at the end of the surrounding function (in this case, main()).
	// The function listener.Close() closes the TCP listener, freeing up the port 8000 when the program terminates.
	// This ensures that resources are properly released, even if an error occurs later in the program.
	defer listener.Close()
	fmt.Printf("server listening on 8000")

	// accepting incoming conn
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("error in connection: %s", err)
			continue
		}
		go handle_conn(conn)
	}
}

// func to handle conn and read data
func handle_conn(conn net.Conn) {
	defer conn.Close()
	// 	bufio.NewReader(conn) initializes a buffered reader that wraps around the conn object (which is of type net.Conn).
	// bufio.Reader provides convenient methods for reading input efficiently, such as ReadString, ReadBytes, or ReadLine.
	// Instead of reading raw bytes manually, the reader allows for line-based or buffered reading.
	reader := bufio.NewReader(conn)
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("error in reading string: %s", err)
			return
		}
		fmt.Printf("message recieved from client: %s", str)
		conn.Write([]byte("server ka reply\n"))
	}
}
