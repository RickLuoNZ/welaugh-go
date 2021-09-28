package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncryption(t *testing.T) {
	e := Encryptor{}
	hash := e.DoEncryption("abc")
	assert.NotEmpty(t, hash)
	assert.NotEqual(t, "abc", hash)
}

func TestPasswordVerification(t *testing.T) {
	e := Encryptor{}
	enPassword := e.DoEncryption("abc")
	assert.True(t, e.VerifyPassword(enPassword, "abc"))
	assert.False(t, e.VerifyPassword(enPassword, "abcd"))
}
