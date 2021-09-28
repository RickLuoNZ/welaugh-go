package model

type SignUpReponse struct {
	Token             string
	TokenExpiresInSec uint64
}

type LoginReponse struct {
	Token             string
	TokenExpiresInSec uint64
}
