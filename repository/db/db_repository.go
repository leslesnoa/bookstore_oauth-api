package db

import (
	"github.com/gocql/gocql"
	"github.com/leslesnoa/bookstore_oauth-api/client/cassandra"
	"github.com/leslesnoa/bookstore_oauth-api/domain/access_token"
	restErr "github.com/leslesnoa/bookstore_oauth-api/utils/errors"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES (?, ?, ?, ?);"
	queryUpdateExpires     = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *restErr.RestErr)
	Create(access_token.AccessToken) *restErr.RestErr
	UpdateExpirationTime(access_token.AccessToken) *restErr.RestErr
}

type dbRepository struct {
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *restErr.RestErr) {
	//TODO: implement get access token from CassandraDB.
	session, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	var result access_token.AccessToken
	if err := session.Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserId,
		&result.ClientId,
		&result.Expires,
	); err != nil {
		if err == gocql.ErrNotFound {
			return nil, restErr.NewNotFoundError("no access token found with given id")
		}
		return nil, restErr.NewInternalServerError(err.Error())
	}

	return &result, nil
	// return nil, restErr.NewInternalServerError("database connection not implemented yet")
	// return nil, nil
}

func (r *dbRepository) Create(at access_token.AccessToken) *restErr.RestErr {
	session, err := cassandra.GetSession()
	if err != nil {
		return restErr.NewInternalServerError(err.Error())
	}
	defer session.Close()

	if err := session.Query(queryCreateAccessToken,
		at.AccessToken,
		at.UserId,
		at.ClientId,
		at.Expires,
	).Exec(); err != nil {
		return restErr.NewInternalServerError(err.Error())
	}
	return nil
}

func (r *dbRepository) UpdateExpirationTime(at access_token.AccessToken) *restErr.RestErr {
	session, err := cassandra.GetSession()
	if err != nil {
		return restErr.NewInternalServerError(err.Error())
	}
	defer session.Close()

	if err := session.Query(queryUpdateExpires,
		at.Expires,
		at.AccessToken,
	).Exec(); err != nil {
		return restErr.NewInternalServerError(err.Error())
	}
	return nil
}
