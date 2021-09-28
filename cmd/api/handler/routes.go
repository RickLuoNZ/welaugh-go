package handler

import (
	"github.com/labstack/echo"
)

func (h *Handler) Register(route *echo.Group) {
	guestUsers := route.Group("", )
	guestUsers.POST("/signup", h.SignUp)
	guestUsers.POST("/login", h.LoginWithEmail)
	guestUsers.GET("/ping", h.Ping)
}