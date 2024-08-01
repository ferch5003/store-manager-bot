package handler

import (
	"backend/internal/platform/vertex"
	"github.com/gofiber/contrib/websocket"
	"log"
)

type HistoryHandler struct {
	vertexClient *vertex.Client
}

func NewHistoryHandler(vertexClient *vertex.Client) *HistoryHandler {
	return &HistoryHandler{
		vertexClient: vertexClient,
	}
}

func (h *HistoryHandler) HistoryChat(c *websocket.Conn) {
	log.Println(c.Locals("allowed"))  // true
	log.Println(c.Params("id"))       // 123
	log.Println(c.Query("v"))         // 1.0
	log.Println(c.Cookies("session")) // ""

	var (
		mt  int
		msg []byte
		err error
	)
	for {
		if mt, msg, err = c.ReadMessage(); err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", msg)

		if err = c.WriteMessage(mt, msg); err != nil {
			log.Println("write:", err)
			break
		}
	}
}
