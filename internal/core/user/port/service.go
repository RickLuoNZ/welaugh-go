package port

import (
	"github.com/rickluonz/pawsitive/internal/core/user/model"
)

type AuthService interface {
	SignUpWithEmail(r model.SignUpRequest)
	AuthLoginWithEmail(r model.LoginRequest)
}