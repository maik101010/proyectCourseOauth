package db

import (
	"errors"
	"fmt"
	"github.com/gocql/gocql"
	"github.com/maik101010/proyectCourseOauth/client/cassandra"
	"github.com/maik101010/proyectCourseOauth/src/domain/access_token"
	"github.com/maik101010/proyectCourseUtilsGoLibrary/rest_errors"
)

const (
	QueryGetAccessToken = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	QueryInsertAccessToken = "INSERT INTO access_tokens (access_token, user_id, client_id, expires) VALUES (?,?,?,?);"
	QueryUpdateAccessToken = "UPDATE access_tokens SET expires=? WHERE access_token=?;"

)
func NewRepository() DbRepository  {
	return &dbRepository{}
}
type DbRepository interface {
	GetById(string) (*access_token.AccessToken, rest_errors.RestError)
	Create(access_token.AccessToken) rest_errors.RestError
	UpdateExpirationTime(access_token.AccessToken) rest_errors.RestError
}
type dbRepository struct {

}
func (r *dbRepository) GetById(id string) (*access_token.AccessToken, rest_errors.RestError){
	var result access_token.AccessToken
	if err := cassandra.GetSession().Query(QueryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserID,
		&result.ClientID,
		&result.Expires,
		); err!=nil{
		if err==gocql.ErrNotFound {
			return nil, rest_errors.NewNotFoundError("no access token found with given id")
		}
		fmt.Println(err)
		return nil, rest_errors.NewInternalServerError("error when trying to get current id", errors.New("database error"))
	}
	return &result, nil
}
func (r * dbRepository) Create(at access_token.AccessToken) rest_errors.RestError {
	if err:= cassandra.GetSession().Query(QueryInsertAccessToken,
		at.AccessToken,
		at.UserID,
		at.ClientID,
		at.Expires).Exec(); err!=nil {
		return rest_errors.NewInternalServerError("error when trying to save access token in database", err)
	}
	return nil
}
func (r * dbRepository) UpdateExpirationTime(token access_token.AccessToken) rest_errors.RestError{
	if err:= cassandra.GetSession().Query(QueryUpdateAccessToken,
		token.Expires,
		token.AccessToken,
		).Exec(); err!=nil {
		return rest_errors.NewInternalServerError("error when trying to update current resource", errors.New("database error"))
	}
	return nil
}
