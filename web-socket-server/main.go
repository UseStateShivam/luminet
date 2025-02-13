// WebSocket multiplexing allows a single WebSocket connection to handle multiple logical communication channels (streams) instead of opening a separate connection for each client/request.
// One WebSocket pipe carries multiple conversations → each identified by a streamId.
// 1️⃣ Accept incoming WebSocket connections from public clients.
// 2️⃣ Assign a unique stream ID for each request (or accept client-sent stream IDs).
// 3️⃣ Forward messages to the localhost backend server.
// 4️⃣ Wait for a response and send it back to the correct client using the streamId.
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

func main() {

}