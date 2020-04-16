package rest

import (
	"encoding/json"
	"errors"
	"github.com/maik101010/proyectCourseOauth/src/domain/users"
	"github.com/maik101010/proyectCourseUtilsGoLibrary/rest_errors"
	"github.com/mercadolibre/golang-restclient/rest"
	"time"
)
var (
	usersRestClient = rest.RequestBuilder{
		BaseURL:        "http://localhost:8081",
		Timeout:100*time.Millisecond,
	}
)
type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, rest_errors.RestError)
}
type usersRepository struct {}

func NewRepository() RestUsersRepository{
	return &usersRepository{}
}
func (r * usersRepository) LoginUser(email string, password string) (*users.User, rest_errors.RestError)  {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}

	response := usersRestClient.Post("/users/login", request)

	if response == nil || response.Response == nil {
		return nil, rest_errors.NewInternalServerError("invalid restclient response when trying to login user", errors.New("restclient error"))
	}

	if response.StatusCode > 299 {
		apiErr, err := rest_errors.NewRestErrorFromBytes(response.Bytes())
		if err != nil {
			return nil, rest_errors.NewInternalServerError("invalid error interface when trying to login user", err)
		}
		return nil, apiErr
	}

	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to unmarshal users login response", errors.New("json parsing error"))
	}
	return &user, nil


}