package hubs

import (
	"bytes"
	"github.com/fasthttp/websocket"
	"log"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 65535
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

// Client is a middleman between the socket connection and the hub.
type Client struct {
	Id  string
	Hub *Hub

	// The socket connection.
	Conn *websocket.Conn

	// Buffered channel of outbound messages.
	Send chan []byte

	MaxMessageSize int64
}

// ReadPump pumps messages from the socket connection to the Hub.
//
// The application runs ReadPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) ReadPump() {
	defer func() {
		c.Hub.unregister <- c
	}()
	c.Conn.SetReadLimit(c.MaxMessageSize)
	_ = c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error { _ = c.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.Hub.broadcast <- message
	}

}

// WritePump pumps messages from the Hub to the socket connection.
//
// A goroutine running WritePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		//_ = c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			_ = c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				if c.Conn != nil {
					_ = c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				}
				return
			}
			if c.Conn != nil {
				w, err := c.Conn.NextWriter(websocket.TextMessage)
				if err != nil {
					return
				}
				_, _ = w.Write(message)

				// Add queued chat messages to the current socket message.
				n := len(c.Send)
				for i := 0; i < n; i++ {
					_, _ = w.Write(newline)
					_, _ = w.Write(<-c.Send)
				}

				if err := w.Close(); err != nil {
					return
				}
			}
		case <-ticker.C:
			if c.Conn != nil {
				_ = c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
				if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
					return
				}
			}
		}
	}
}
