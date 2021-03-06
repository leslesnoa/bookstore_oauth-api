package access_token

import (
	"strings"

	"github.com/leslesnoa/bookstore_oauth-api/domain/access_token"
	"github.com/leslesnoa/bookstore_oauth-api/repository/db"
	"github.com/leslesnoa/bookstore_oauth-api/repository/rest"
	restErr "github.com/leslesnoa/bookstore_oauth-api/utils/errors"
)

type Service interface {
	GetById(string) (*access_token.AccessToken, *restErr.RestErr)
	Create(access_token.AccessTokenRequest) (*access_token.AccessToken, *restErr.RestErr)
	UpdateExpirationTime(access_token.AccessToken) *restErr.RestErr
}

type service struct {
	restUsersRepo rest.RestUsersRepository
	dbRepo        db.DbRepository
}

func NewService(usersRepo rest.RestUsersRepository, dbRepo db.DbRepository) Service {
	return &service{
		restUsersRepo: usersRepo,
		dbRepo:        dbRepo,
	}
}

func (s *service) GetById(accessTokenId string) (*access_token.AccessToken, *restErr.RestErr) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, restErr.NewBadRequestError("invalid access token id")
	}
	accessToken, err := s.dbRepo.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(request access_token.AccessTokenRequest) (*access_token.AccessToken, *restErr.RestErr) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	//TODO: Support both grant types: client_credentials and password

	// Authenticate the user against the Users API:
	user, err := s.restUsersRepo.LoginUser(request.Username, request.Password)
	if err != nil {
		return nil, err
	}

	// Generate a new access token:
	at := access_token.GetNewAccessToken(user.Id)
	at.Generate()

	// Save the new access token in Cassandra:
	if err := s.dbRepo.Create(at); err != nil {
		return nil, err
	}
	return &at, nil
}

func (s *service) UpdateExpirationTime(at access_token.AccessToken) *restErr.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.dbRepo.UpdateExpirationTime(at)
}
