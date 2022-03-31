package main

import (
	"github.com/gorilla/websocket"
)

// clientはチャットを行っている1人のユーザーを表す
type client struct {
	// socketはこのクライアントのためのWebSocket
	socket *websocket.Conn
	// sendはメッセージが送られるチャネル
	send chan []byte
	// roomはこのクライアントが参加しているチャットルーム
	room *room
}

func (c *client) read() {
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			break
		}
		c.room.forward <- msg
	}

	c.socket.Close()
}

func (c *client) write() {
	for msg := range c.send {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			break
		}
	}

	c.socket.Close()
}
