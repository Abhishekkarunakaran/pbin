package ports

import (
	"github.com/labstack/echo/v4"
)

type Handler interface {
	PasteData(e echo.Context) error
	GetData(e echo.Context) error
}
