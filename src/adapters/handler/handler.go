package handler

import (
	"net/http"

	log "log/slog"

	"github.com/Abhishekkarunakaran/pbin/src/core/domain"
	"github.com/Abhishekkarunakaran/pbin/src/core/ports"
	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
)

type handler struct {
	service ports.Service
}

// GetData implements ports.Handler.
func (h *handler) GetData(e echo.Context) error {
	panic("unimplemented")
}

func NewHandler(service ports.Service) ports.Handler {
	return &handler{
		service: service,
	}
}

// PasteData implements ports.Handler.
func (h *handler) PasteData(e echo.Context) error {
	ctx := e.Request().Context()

	var payload domain.Payload
	if err := e.Bind(&payload); err != nil {
		log.Error(err.Error())
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	if err := payload.ValidErr(); err != nil {
		log.Error(err.Error())
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	id, err := h.service.SaveContent(ctx, &payload)
	if err != nil {
		log.Error(err.Error())
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, map[string]uuid.UUID{
		"id": id,
	})

}
