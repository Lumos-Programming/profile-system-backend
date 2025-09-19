package jwt

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestJWT_IssueAndAuthenticate(t *testing.T) {
	claims := jwt.RegisteredClaims{
		Subject:   "TestSubject",
		Issuer:    "TestIssuer",
		IssuedAt:  jwt.NewNumericDate(time.Now().Add(-5 * time.Hour)),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	}
	secret := "testSecret"

	// JWT発行
	tokenString, err := IssueJWT(claims, secret)
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)

	// 正しいシークレットで認証
	token, err := AuthenticateJWT(tokenString, []byte(secret))
	assert.NoError(t, err)
	assert.True(t, token.Valid)

	// 不正なシークレットで認証（失敗するはず）
	token2, err2 := AuthenticateJWT(tokenString, []byte("wrongSecret"))
	assert.Error(t, err2)
	assert.Nil(t, token2)
}

func TestCreateClaims(t *testing.T) {
	tests := []struct {
		name   string
		id     string
		secret string
	}{
		{
			name:   "test1",
			id:     "testing123",
			secret: "ThisIsSecret",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			claims := CreateClaims(tt.id, 24*time.Hour, "testIssuer")
			got, err := IssueJWT(claims, tt.secret)
			if err != nil {
				t.Errorf("IssueJWT() error = %v", err)
				return
			}
			verify, err := Verify(got, tt.secret)
			assert.NoError(t, err)
			assert.Equal(t, claims.Id, verify.Id)
		})
	}
}
