package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/rickluonz/pawsitive/cmd/api/db"
	"github.com/rickluonz/pawsitive/cmd/api/repository/model"
	"github.com/rickluonz/pawsitive/internal/core/user/entity"
)

type UserRepository struct {
	DB *db.DB
}

func (r *UserRepository) Save(user entity.User) (entity.User, error) {
	newUser := model.User{Email: user.Email, Password: user.Password}
	if err := r.DB.Client.Create(&newUser).Error; err != nil {
		return entity.User{}, err
	}

	return r.GetUserByEmail(user.Email)
}

func (r *UserRepository) GetUserByEmail(email string) (entity.User, error) {

	var m model.User
	if err := r.DB.Client.Where(&model.User{Email: email}).First(&m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return entity.User{}, nil
		}
		return entity.User{}, err
	}

	return entity.User{
		Id: m.Id, 
		Email: m.Email,
		Password: m.Password,
		}, nil
}
