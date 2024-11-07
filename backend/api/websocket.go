package api

import (
	"job-scheduler-backend/broadcast"
	"github.com/gofiber/websocket/v2"
)

// HandleWebSocket manages WebSocket connections.
func HandleWebSocket(c *websocket.Conn) {
	broadcast.RegisterClient(c)

	defer func() {
		broadcast.UnregisterClient(c)
	}()

	for {
		_, _, err := c.ReadMessage()
		if err != nil {
			break
		}
	}
}
