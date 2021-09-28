package port

import (
	"github.com/rickluonz/pawsitive/internal/core/user/model"
)

type AuthPresenter interface {
	CreateAuthResponse(r model.SignUpReponse) interface{}
}