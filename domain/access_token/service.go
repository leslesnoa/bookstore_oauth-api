package access_token

import (
	"strings"

	restErr "github.com/leslesnoa/bookstore_oauth-api/utils/errors"
)

type Repository interface {
	GetById(string) (*AccessToken, *restErr.RestErr)
	Create(AccessToken) *restErr.RestErr
	UpdateExpirationTime(AccessToken) *restErr.RestErr
}

type Service interface {
	GetById(string) (*AccessToken, *restErr.RestErr)
	Create(AccessToken) *restErr.RestErr
	UpdateExpirationTime(AccessToken) *restErr.RestErr
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

func (s *service) Create(at AccessToken) *restErr.RestErr {
	// Validate()実装前
	// at.AccessToken = strings.TrimSpace(at.AccessToken)
	// if len(at.AccessToken) == 0 {
	// 	return restErr.NewBadRequestError("invalid access token id")
	// }
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.Create(at)
	// return nil
}

func (s *service) UpdateExpirationTime(at AccessToken) *restErr.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.UpdateExpirationTime(at)
	// return nil
}
