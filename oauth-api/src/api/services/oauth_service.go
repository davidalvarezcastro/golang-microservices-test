package services

import (
	"time"

	"github.com/davidalvarezcastro/golang-microservices-test/oauth-api/src/api/model/oauth"
	"github.com/davidalvarezcastro/golang-microservices-test/src/api/utils/errors"
)

type oauthService struct {
}

type oauthServiceInterface interface {
	CreateAccessToken(request oauth.AccessTokenRequest) (*oauth.AccessToken, errors.APIError)
	GetAccessToken(accessToken string) (*oauth.AccessToken, errors.APIError)
}

var (
	// OauthService is the interface of our oauth service
	OauthService oauthServiceInterface
)

func init() {
	OauthService = &oauthService{}
}

func (s *oauthService) CreateAccessToken(request oauth.AccessTokenRequest) (*oauth.AccessToken, errors.APIError) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	user, err := oauth.GetUserByUsernameAndPassword(request.Username, request.Password)
	if err != nil {
		return nil, err
	}

	token := oauth.AccessToken{
		UserID:  user.ID,
		Expires: time.Now().UTC().Add(24 * time.Hour).Unix(),
	}
	if err := token.Save(); err != nil {
		return nil, err
	}

	return &token, nil
}

func (s *oauthService) GetAccessToken(accessToken string) (*oauth.AccessToken, errors.APIError) {
	token, err := oauth.GetAccessTokenByToken(accessToken)
	if err != nil {
		return nil, err
	}

	if token.IsExpired() {
		return nil, errors.NewNotFoundAPIError("no access token found with given parameters")
	}

	return token, err
}
