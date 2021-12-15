package rest

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/leslesnoa/bookstore_oauth-api/domain/users"
	restErr "github.com/leslesnoa/bookstore_oauth-api/utils/errors"
	"github.com/mercadolibre/golang-restclient/rest"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "https://api.bookstore.com",
		Timeout: 100 * time.Millisecond,
	}
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *restErr.RestErr)
}

type usersRepository struct{}

func NewRepository() RestUsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email string, password string) (*users.User, *restErr.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}

	// debug for request statement
	bytes, _ := json.Marshal(request)
	fmt.Println(string(bytes))

	response := usersRestClient.Post("/users/login", request)
	fmt.Printf("response %s", response)
	fmt.Println(response.StatusCode)
	if response == nil || response.Response == nil {
		return nil, restErr.NewInternalServerError("invalid restclient response when trying to login user")
	}
	if response.StatusCode > 299 {
		fmt.Println(response.String())
		var resterr restErr.RestErr
		err := json.Unmarshal(response.Bytes(), &resterr)
		if err != nil {
			return nil, restErr.NewInternalServerError("invalid error interface when trying to login user")
		}
	}
	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, restErr.NewInternalServerError("error when trying to unmarshal users response")
	}
	return &user, nil
}
