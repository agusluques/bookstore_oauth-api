package access_token

import (
	"fmt"
	"strings"
	"time"

	"github.com/agusluques/bookstore_oauth-api/src/utils/crypto_utils"
	"github.com/agusluques/bookstore_oauth-api/src/utils/errors"
)

const (
	expirationTime             = 24
	grantTypePassword          = "password"
	grantTypeClientCredentials = "client_credentials"
)

// AccessTokenRequest object
type AccessTokenRequest struct {
	GrantType string `json:"grant_type" binding:"required"`
	Scope     string `json:"scope"`

	// password grant type
	Username string `json:"username"`
	Password string `json:"password"`

	// client credentials grant type
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

// Validate validates an access token
func (at *AccessTokenRequest) Validate() *errors.RestError {
	if at.GrantType != grantTypePassword && at.GrantType != grantTypeClientCredentials {
		return errors.NewBadRequestError("invalid grant type")
	}
	return nil
}

// AccessToken object
type AccessToken struct {
	AccessToken string `json:"access_token" binding:"required"`
	UserID      int64  `json:"user_id" binding:"required"`
	ClientID    int64  `json:"client_id" binding:"required"`
	Expires     int64  `json:"expires" binding:"required"`
}

// Validate validates an access token
func (at *AccessToken) Validate() *errors.RestError {
	at.AccessToken = strings.TrimSpace(at.AccessToken)

	if len(at.AccessToken) == 0 {
		return errors.NewBadRequestError("invalid access token id")
	}
	if at.UserID <= 0 {
		return errors.NewBadRequestError("invalid user id")
	}
	if at.ClientID <= 0 {
		return errors.NewBadRequestError("invalid client id")
	}
	if at.Expires <= 0 {
		return errors.NewBadRequestError("invalid expiration time")
	}
	return nil
}

// GetNewAccessToken returns new access token
func GetNewAccessToken(userID int64) AccessToken {
	return AccessToken{
		UserID:  userID,
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) isExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}

func (at *AccessToken) Generate() {
	at.AccessToken = crypto_utils.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserID, at.Expires))
}
