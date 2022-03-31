package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "chat.html", gin.H{})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	
	room := newRoom()
	r.GET("/room", func(c *gin.Context) {
		room.ServeHTTP(c.Writer, c.Request)
	})

	// チャットルームの開始
	go room.run()

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	err := r.Run(":8080")

	if err != nil {
		log.Fatal("gin Run:", err)
	}
}
