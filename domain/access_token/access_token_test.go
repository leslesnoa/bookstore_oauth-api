package access_token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstants(t *testing.T) {
	// if expirationTime != 24 {
	// 	t.Error("expiration time should be 24 hours")
	// }
	assert.EqualValues(t, 24, expirationTime, "expiration time should be 24 hours")
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken(1)
	// if at.IsExpired() {
	// 	t.Error("brand new access token should not be expired")
	// }
	assert.False(t, at.IsExpired(), "brand new access token should not be expired")

	// if at.AccessToken != "" {
	// 	t.Error("new access token should not have defined access token id")
	// }
	assert.EqualValues(t, "", at.AccessToken, "new access token should not have defined access token id")

	// if at.UserId != 0 {
	// 	t.Error("new access token should not have an associated user id")
	// }
	assert.True(t, at.UserId == 0, "new access token should not have an associated user id")
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	// if !at.IsExpired() {
	// 	t.Error("empty access token should be expired by default")
	// }
	assert.True(t, at.IsExpired(), "empty access token should be expired by default")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	// if at.IsExpired() {
	// 	t.Error("access token expiring three hours from now should NOT be expired")
	// }
	assert.False(t, at.IsExpired(), "access token expiring three hours from now should NOT be expired")
}
