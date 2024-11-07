package broadcast

import (
	"sync"
	"github.com/gofiber/websocket/v2"
	"job-scheduler-backend/types"
)

var (
	clients = make(map[*websocket.Conn]bool)
	wsMutex sync.Mutex
)

// RegisterClient adds a WebSocket client to the broadcast list.
func RegisterClient(c *websocket.Conn) {
	wsMutex.Lock()
	clients[c] = true
	wsMutex.Unlock()
}

// UnregisterClient removes a WebSocket client from the broadcast list.
func UnregisterClient(c *websocket.Conn) {
	wsMutex.Lock()
	delete(clients, c)
	wsMutex.Unlock()
	c.Close()
}

// UpdateClients broadcasts job updates to all WebSocket clients.
func UpdateClients(jobs []types.Job) {
	wsMutex.Lock()
	defer wsMutex.Unlock()

	for client := range clients {
		client.WriteJSON(jobs)
	}
}
