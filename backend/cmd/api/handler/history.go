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
	"log"
)

type HistoryHandler struct {
	vertexClient       *vertex.Client
	historyService     history.Service
	defaultBotResponse []byte
}

var _defaultBotResponse = domain.History{
	BotResponse: "No se ha podido procesar una respuesta en concreto, puede que haya un problema con el mensaje o el" +
		"serviceo, vuelva a intentarlo mas tarde.",
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

	vertexConfig := vertex.Config{
		ProjectTunedID:  configurations.ProjectTunedID,
		RegionTuned:     configurations.RegionTuned,
		EndpointTunedID: configurations.EndpointTunedID,

		ProjectFlashID: configurations.ProjectFlashID,
		RegionFlash:    configurations.RegionFlash,
		ModelFlashID:   configurations.ModelFlashID,
	}
	vertexClient, err := vertex.NewVertexClient(
		ctx,
		vertexConfig,
		credentials,
		baseHistories,
	)
	if err != nil {
		return nil, err
	}

	botMsg, err := json.Marshal(_defaultBotResponse)
	if err != nil {
		return nil, err
	}

	return &HistoryHandler{
		vertexClient:       vertexClient,
		historyService:     historyService,
		defaultBotResponse: botMsg,
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

		log.Printf("USER MESSAGE: %+v", newHistory)

		if newHistory.UserMessage == "" {
			break
		}

		botResponse, err := h.vertexClient.SendMessage(newHistory.UserMessage)
		if err != nil {
			if err = c.WriteMessage(mt, h.defaultBotResponse); err != nil {
				break
			}
		}

		newHistory.BotResponse = botResponse
		if _, isMIMEContent := files.GetMIMEType(newHistory.UserMessage); isMIMEContent {
			newHistory.Multimedia = true
		}

		var savedHistory domain.History
		if savedHistory, err = h.historyService.Save(context.Background(), newHistory); err != nil {
			break
		}

		log.Printf("BOT RESPONSE %+v", savedHistory)

		botMsg, err := json.Marshal(savedHistory)
		if err != nil {
			break
		}

		if err = c.WriteMessage(mt, botMsg); err != nil {
			break
		}
	}
}
