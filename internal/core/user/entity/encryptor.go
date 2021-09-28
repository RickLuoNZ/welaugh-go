package entity

import (
	"golang.org/x/crypto/bcrypt"
)

type Encryptor struct{}

// encrypt a string with bcrypt
func (e *Encryptor) DoEncryption(s string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
    return string(hash);
}

// verify a user's password
func (e *Encryptor) VerifyPassword(userPasswd, inputPasswd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(userPasswd), []byte(inputPasswd))
	return err == nil
}