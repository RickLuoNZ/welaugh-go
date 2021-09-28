package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/rickluonz/pawsitive/cmd/api/repository"
	"github.com/rickluonz/pawsitive/internal/core/user/model"
	"github.com/rickluonz/pawsitive/internal/core/user/service"
	"github.com/rickluonz/pawsitive/pkg/logger"
)

type signUpRequest struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (r *signUpRequest) bind(c echo.Context, u *model.SignUpRequest) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	u.Email = r.Login
	u.Password = r.Password
	return nil
}

type authResponse struct {
	Token             string `json:"token"`
	TokenExpiresInSec uint64 `json:"token_expires_in_sec"`
}

func (h *Handler) SignUp(c echo.Context) error {
	signUpReq := model.SignUpRequest{}
	httpReq := &signUpRequest{}
	if err := httpReq.bind(c, &signUpReq); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	userRepo := repository.UserRepository{DB: h.DB}
	authService := service.AuthService{Repository: &userRepo}
	
	authResult, err := authService.SignUpWithEmail(signUpReq)
	if err != nil {
		logger.Info.Println("auth" + err.Error())
		return c.JSON(http.StatusUnprocessableEntity, err)
	}
	
	return c.JSON(http.StatusCreated, buildSignUpResponse(authResult))
}

func buildSignUpResponse(result model.SignUpReponse) *authResponse {
	r := new(authResponse)
	r.Token = result.Token
	r.TokenExpiresInSec = result.TokenExpiresInSec
	return r
}

type userLoginWithEmailRequest struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (r *userLoginWithEmailRequest) bind(c echo.Context, u *model.LoginRequest) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}

	u.Email = r.Login
	u.Password = r.Password
	return nil
}

type loginResponse struct {
	Token             string `json:"token"`
	TokenExpiresInSec uint64 `json:"token_expires_in_sec"`
}

func (h *Handler) LoginWithEmail(c echo.Context) error {
	loginReq := model.LoginRequest{}
	httpReq := &userLoginWithEmailRequest{}
	if err := httpReq.bind(c, &loginReq); err != nil {
		//c.Logger().Debug(err)
		return c.JSON(http.StatusUnprocessableEntity, err)
	}
	//c.Logger().Debug("login:" + loginReq.Email)
	//c.Logger().Debug("password:" + loginReq.Password)
	userRepo := repository.UserRepository{DB: h.DB}
	authService := service.AuthService{Repository: &userRepo}
	
	loginReponse, err := authService.AuthLoginWithEmail(loginReq)
	if err != nil {
		logger.Info.Println("auth error: " + err.Error())
		return c.JSON(http.StatusUnprocessableEntity, err)
	}
	
	return c.JSON(http.StatusOK, buildLoginResponse(loginReponse))
}

func buildLoginResponse(result model.LoginReponse) *loginResponse {
	r := new(loginResponse)
	r.Token = result.Token
	r.TokenExpiresInSec = result.TokenExpiresInSec
	return r
}

