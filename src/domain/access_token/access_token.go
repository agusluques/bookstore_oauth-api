package access_token

import (
	"time"
)

const expirationTime = 24

// AccessToken object
type AccessToken struct {
	AccessToken string `json:access_token`
	UserID      int64  `json:user_id`
	ClientID    int64  `json:client_id`
	Expires     int64  `json:expires`
}

// GetNewAccessToken returns new access token
func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) isExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}
