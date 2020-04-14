package access_token

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAccessTokenConstants(t *testing.T){
	assert.EqualValues(t, 24, expirationTime, "expiration time should be 24 hours ")
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	assert.False(t, at.isExpired(), "Brand new access token should not be expired")
	//assert.EqualValues(t, at.AccessToken, "new access token should not have defined access token id")
	assert.True(t, at.UserID==0, "new access token should not have an associated user id")
}
func TestAccessTokenIsExpired(t *testing.T){
	at := AccessToken{}
	assert.True(t, at.isExpired(), "empty acces token should be expired by default")
	if !at.isExpired() {
		t.Error("")
	}
	at.Expires = time.Now().UTC().Add(3*time.Hour).Unix()
	assert.False(t, at.isExpired(), "access token expiring three hours from now should NOT be expired")
}
