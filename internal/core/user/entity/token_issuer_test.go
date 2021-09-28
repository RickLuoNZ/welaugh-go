package entity

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenCreationWithDefaultExp(t *testing.T) {
	i := TokenIssuer{}
	token := i.NewToken(1)
	assert.NotEmpty(t, token.Value)

	expInDays, _ := strconv.ParseUint(TOKEN_EXPIRES_IN_DAY, 10, 64)
	expInSeconds := 24 * 60 * 60 * expInDays
	assert.Equal(t, expInSeconds, token.ExpiresInSec)
}

func TestTokenCreationWithConfigExp(t *testing.T) {
	os.Setenv("TOKEN_EXPIRES_IN_DAY", "10")
	i := TokenIssuer{}
	token := i.NewToken(1)
	assert.NotEmpty(t, token.Value)

	expInSeconds := uint64(24 * 60 * 60 * 10)
	assert.Equal(t, expInSeconds, token.ExpiresInSec)
}