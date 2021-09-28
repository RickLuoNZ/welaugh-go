package port

import (
	"github.com/rickluonz/pawsitive/internal/core/user/entity"
)

type UserRepository interface {
	Save(user entity.User) (entity.User, error)
	GetUserByEmail(email string) (entity.User, error)
}