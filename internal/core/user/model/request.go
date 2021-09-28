package model

type SignUpRequest struct {
	Username    string
	Email       string
	Password    string
	PhoneNumber string
}

type LoginRequest struct {
	Email       string
	PhoneNumber string
	Password    string
}
