package access_token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstant(t *testing.T) {
	assert.Equal(t, 24, expirationTime, "expiration time should be 24 hours")
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken(0)

	assert.False(t, at.isExpired(), "new access token should not be expired")
	assert.Equal(t, "", at.AccessToken, "new access token should not have access token id")
	assert.EqualValues(t, 0, at.UserID, "new access token should not have asociated user id")
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}

	assert.True(t, at.isExpired(), "access token should be expired by default")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.isExpired(), "access token created 3 hours from now should not be expired")
}
