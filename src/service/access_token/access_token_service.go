package access_token

import (
	"github.com/maik101010/proyectCourseOauth/src/domain/access_token"
	"github.com/maik101010/proyectCourseOauth/src/repository/db"
	"github.com/maik101010/proyectCourseOauth/src/repository/rest"
	"github.com/maik101010/proyectCourseUtilsGoLibrary/rest_errors"
	"strings"
)

type Service interface {
	GetById(string) (*access_token.AccessToken, rest_errors.RestError)
	Create(access_token.AccessTokenRequest) (*access_token.AccessToken, rest_errors.RestError)
	UpdateExpirationTime(access_token.AccessToken) rest_errors.RestError
}
type service struct {
	restUsersRepo rest.RestUsersRepository
	dbRepo        db.DbRepository
}
func NewService(usersRepo rest.RestUsersRepository, repoDB db.DbRepository) Service{
	return &service{
		restUsersRepo: usersRepo,
		dbRepo:        repoDB,
	}
}
func (s *service)GetById(accessTokenId string) (*access_token.AccessToken, rest_errors.RestError){
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId)==0 {
		return nil, rest_errors.NewBadRequestError("invalid access token id")
	}
	accessToken, err := s.dbRepo.GetById(accessTokenId)
	if err!=nil {
		return nil, err
	}
	return accessToken, nil
}
func (s *service) Create(atr access_token.AccessTokenRequest) (*access_token.AccessToken, rest_errors.RestError) {
	if err := atr.Validate(); err!=nil{
		return nil, err
	}
	//TODO: Support both grant types: client_credentials and password
	//Authenticate the user against the Users Api:
	user, err := s.restUsersRepo.LoginUser(atr.UserName, atr.Password)
	if err!=nil {
		return nil, err
	}
	at := access_token.GetNewAccessToken(user.Id)
	at.Generate()
	if err:=s.dbRepo.Create(at); err!=nil {
		return nil, err
	}
	return &at, nil
}

func (s *service) UpdateExpirationTime(at access_token.AccessToken)  rest_errors.RestError {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.dbRepo.UpdateExpirationTime(at)
}
