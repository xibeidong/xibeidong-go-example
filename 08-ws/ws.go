package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

func main() {
	gin.SetMode(gin.DebugMode) //debug
	engine := gin.Default()

	engine.GET("/ws", handlerMobile)
	if err := engine.Run(":8001"); err != nil {
		log.Println(err)
		return
	}
}
func handlerMobile(ctx *gin.Context) {
	//升级get请求为webSocket协议

	ws, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return
	}

	for {
		_, bytes, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		m := message{}
		err = json.Unmarshal(bytes, &m)
		if err != nil {
			log.Println(err)
		} else {
			log.Println(m.Data)
		}

		m.Data = time.Now().Format("2006-01-02 15:04:05")
		m.Type = "server-answer"
		ws.WriteJSON(&m)
	}

}
