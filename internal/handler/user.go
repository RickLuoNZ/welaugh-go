package handler

import (
	"github.com/rickluonz/pawsitive/internal/core/user/model"
	"github.com/rickluonz/pawsitive/internal/core/user/port"
	"github.com/rickluonz/pawsitive/internal/core/user/service"
)

type AuthHandler struct {
	authService  port.AuthService
}

func (h *AuthHandler) SignUpWithEmail() error {
	signUpRequest := model.SignUpRequest{}
    authSrv := service.AuthService{}
	authSrv.SignUpWithEmail(signUpRequest)
	return nil
}