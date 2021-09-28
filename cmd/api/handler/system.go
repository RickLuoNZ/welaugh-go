package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func (h *Handler) Ping(c echo.Context) error {
	return c.JSON(http.StatusCreated, "pong")
}
