package service

import (
	"errors"
	"github.com/rickluonz/pawsitive/internal/core/user/entity"
	"github.com/rickluonz/pawsitive/internal/core/user/model"
	"github.com/rickluonz/pawsitive/internal/core/user/port"
)

type AuthService struct {
	Repository port.UserRepository
}

// SignUp a new user with email and password
func (srv *AuthService) SignUpWithEmail(signUpRequest model.SignUpRequest) (model.SignUpReponse, error) {
	encryptor := entity.Encryptor{}
	encryptedPassword := encryptor.DoEncryption(signUpRequest.Password)

	newUser := entity.User{
		Email:    signUpRequest.Email,
		Password: encryptedPassword,
	}

	response := model.SignUpReponse{}
	savedUser, err := srv.Repository.Save(newUser)
	if err != nil {
		return response, err
	}

	tokenIssuer := entity.TokenIssuer{}
	token := tokenIssuer.NewToken(savedUser.Id)

	response.Token = token.Value
	response.TokenExpiresInSec = token.ExpiresInSec

	return response, nil
}

// Authenticate a user with email and password
func (srv *AuthService) AuthLoginWithEmail(loginRequest model.LoginRequest) (model.LoginReponse, error) {
    response := model.LoginReponse{}

	user, err := srv.Repository.GetUserByEmail(loginRequest.Email)
	if err != nil {
		return response, errors.New("User does not exist!")
	}

	encryptor := entity.Encryptor{}
	if !encryptor.VerifyPassword(user.Password, loginRequest.Password) {
		return response, errors.New("Incorrect password!")
	}

	tokenIssuer := entity.TokenIssuer{}
	token := tokenIssuer.NewToken(user.Id)

	response.Token = token.Value
	response.TokenExpiresInSec = token.ExpiresInSec

	return response, nil
}
