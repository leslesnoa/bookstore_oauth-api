package db

import (
	"github.com/leslesnoa/bookstore_oauth-api/client/cassandra"
	"github.com/leslesnoa/bookstore_oauth-api/domain/access_token"
	restErr "github.com/leslesnoa/bookstore_oauth-api/utils/errors"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *restErr.RestErr)
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
	return nil, restErr.NewInternalServerError("database connection not implemented yet")
	// return nil, nil
}
