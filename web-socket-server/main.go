// WebSocket multiplexing allows a single WebSocket connection to handle multiple logical communication channels (streams) instead of opening a separate connection for each client/request.
// One WebSocket pipe carries multiple conversations → each identified by a streamId.
// 1️⃣ Single WebSocket Connection Handles Multiple Streams → Instead of opening multiple WebSocket connections, you reuse a single connection for multiple conversations.
// 2️⃣ Each "Stream" is Identified by a streamId → Messages are tagged with a unique streamId so the server knows which request they belong to.
// 3️⃣ WebSocket Server Acts as a Router → It receives, processes, and forwards messages between public clients and the backend.
// 4️⃣ Connection Pool Maintains Active Clients → The server stores a mapping of WebSocket connections and their active streams.
// Connection pool to maintain record of active clients
// Pool looks something like:
// {
//     WebSocketConnection1: {
//         streamId1: "UUID-1",
//         streamId2: "UUID-2",
//     },
//     WebSocketConnection2: {
//         streamId3: "UUID-3",
//         streamId4: "UUID-4",
//     }
// }

package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/ws", web_socket_handler)
	fmt.Println("Web socket server started at /ws")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error listening web socket server: %s", err)
		os.Exit(1)
	}
}

func web_socket_handler(w http.ResponseWriter, r *http.Request) {

}