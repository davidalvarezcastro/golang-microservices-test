package oauth

import "time"

// AccessToken stores data about the access token
type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	Expires     int64  `json:"expires"`
}

// IsExpired returns if the token is expired
func (at *AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).UTC().Before(time.Now().UTC())
}
