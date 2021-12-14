package access_token

import (
	"strings"

	restErr "github.com/leslesnoa/bookstore_oauth-api/utils/errors"
)

type Repository interface {
	GetById(string) (*AccessToken, *restErr.RestErr)
}

type Service interface {
	GetById(string) (*AccessToken, *restErr.RestErr)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(accessTokenId string) (*AccessToken, *restErr.RestErr) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, restErr.NewBadRequestError("invalid access token id")
	}

	accessToken, err := s.repository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
	// return s.repository.GetById(accessTokenId)
}
