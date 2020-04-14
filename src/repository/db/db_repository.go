package db

import (
	"fmt"
	"github.com/gocql/gocql"
	"github.com/maik101010/proyectCourseOauth/client/cassandra"
	"github.com/maik101010/proyectCourseOauth/src/domain/access_token"
	"github.com/maik101010/proyectCourseOauth/src/util/errors"
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
	GetById(string) (*access_token.AccessToken, *errors.RestError)
	Create(access_token.AccessToken) *errors.RestError
	UpdateExpirationTime(access_token.AccessToken) *errors.RestError
}
type dbRepository struct {

}
func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestError){
	var result access_token.AccessToken
	if err := cassandra.GetSession().Query(QueryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserID,
		&result.ClientID,
		&result.Expires,
		); err!=nil{
		if err==gocql.ErrNotFound {
			return nil, errors.NewNotFoundError("No access token found with give id")
		}
		fmt.Println(err)
		return nil, errors.NewInternalServerError(err.Error())
	}
	return &result, nil
}
func (r * dbRepository) Create(at access_token.AccessToken) *errors.RestError {
	if err:= cassandra.GetSession().Query(QueryInsertAccessToken,
		at.AccessToken,
		at.UserID,
		at.ClientID,
		at.Expires).Exec(); err!=nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}
func (r * dbRepository) UpdateExpirationTime(token access_token.AccessToken) *errors.RestError{
	if err:= cassandra.GetSession().Query(QueryUpdateAccessToken,
		token.Expires,
		token.AccessToken,
		).Exec(); err!=nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}
