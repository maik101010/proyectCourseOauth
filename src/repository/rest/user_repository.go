package rest

import (
	"encoding/json"
	"fmt"
	"github.com/maik101010/proyectCourseOauth/src/domain/users"
	"github.com/maik101010/proyectCourseOauth/src/util/errors"
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
	LoginUser(string, string) (*users.User, *errors.RestError)
}
type usersRepository struct {}

func NewRepository() RestUsersRepository{
	return &usersRepository{}
}
func (r * usersRepository) LoginUser(email string, password string) (*users.User, *errors.RestError)  {
	request:= users.UserLoginRequest{
		Email:    email,
		Password: password,
	}

	response := usersRestClient.Post("/users/login", request)
	if response ==nil || response.Response==nil{
		return nil, errors.NewInternalServerError("invalid restclient respose when trying to login user")
	}
	if response.StatusCode>299 {
		fmt.Println(response.String())
		var restErr errors.RestError
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err!=nil {
			return nil, errors.NewInternalServerError("invalid error interface when trying to login user")
		}
		return nil, &restErr
	}
	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err!=nil  {
		return nil, errors.NewInternalServerError("error when trying on marschall your users response")
	}
	return &user, nil


}