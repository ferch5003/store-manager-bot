package handler

import (
	"backend/config"
	"backend/internal/domain"
	"backend/internal/history"
	"backend/internal/platform/files"
	"backend/internal/platform/vertex"
	"context"
	"encoding/json"
	"github.com/gofiber/contrib/websocket"
)

type HistoryHandler struct {
	vertexClient   *vertex.Client
	historyService history.Service
}

func NewHistoryHandler(configurations *config.EnvVars, historyService history.Service) (*HistoryHandler, error) {
	credentials, err := files.GetFile("config/client_secret.json")
	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	baseHistories, err := historyService.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	vertexClient, err := vertex.NewVertexClient(
		ctx,
		configurations.ProjectID,
		configurations.Region,
		configurations.EndpointID,
		credentials,
		baseHistories,
	)
	if err != nil {
		return nil, err
	}

	return &HistoryHandler{
		vertexClient:   vertexClient,
		historyService: historyService,
	}, nil
}

func (h *HistoryHandler) HistoryChat(c *websocket.Conn) {
	var (
		mt  int
		msg []byte
		err error
	)
	for {
		if mt, msg, err = c.ReadMessage(); err != nil {
			break
		}

		var newHistory domain.History
		if err = json.Unmarshal(msg, &newHistory); err != nil {
			break
		}

		botResponse, err := h.vertexClient.SendMessage(newHistory.UserMessage)
		if err != nil {
			break
		}

		newHistory.BotResponse = botResponse

		var savedHistory domain.History
		if savedHistory, err = h.historyService.Save(context.Background(), newHistory); err != nil {
			break
		}

		botMsg, err := json.Marshal(savedHistory)
		if err != nil {
			break
		}

		if err = c.WriteMessage(mt, botMsg); err != nil {
			break
		}
	}
}
