package access_token

import (
	"fmt"
	"github.com/maik101010/proyectCourseUtilsGoLibrary/crypto_utils"
	"github.com/maik101010/proyectCourseUtilsGoLibrary/rest_errors"
	"strings"
	"time"
)

const (
	expirationTime = 24
	grantTypePassword = "password"
	grantTypeCredential = "client_credentials"
)

//AccessTokenRequest struct
type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope int64 `json:"scope"`

	//Used for password grant_type
	UserName string `json:"user_name"`
	Password      string  `json:"password"`
	//Used for Client credentials grant type
	ClientId      string  `json:"client_id"`
	ClientSecret      string  `json:"client_secret"`

}
func (at * AccessTokenRequest) Validate() rest_errors.RestError  {
	switch at.GrantType {
	case grantTypePassword:
	case grantTypeCredential:
	default:
		return rest_errors.NewBadRequestError("invalid grant_type parameter")
	}
	return nil
}

//AccessToken struct
type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	ClientID    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func (at * AccessToken) Validate() rest_errors.RestError  {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken=="" {
		return rest_errors.NewBadRequestError("invalid access token id")
	}
	if at.UserID<=0 {
		return rest_errors.NewBadRequestError("invalid user id")
	}
	if at.ClientID<=0 {
		return rest_errors.NewBadRequestError("invalid client id")
	}
	if at.Expires <=0 {
		return rest_errors.NewBadRequestError("invalid expiration id")
	}
	return nil
}
//GetNewAccessToken return struct AccessToken
func GetNewAccessToken(userId int64) AccessToken {
	return AccessToken{
		UserID:userId,
		Expires: time.Now().UTC().Add(expirationTime + time.Hour).Unix(),
	}
}

func (at AccessToken) isExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}
func (at *AccessToken) Generate() {
	at.AccessToken = crypto_utils.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserID, at.Expires))
}