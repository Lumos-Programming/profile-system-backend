package jwt

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

type Manager struct {
	secret []byte `json:"student_id"`
}

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

// JWTトークンを認証するだけの関数例
func (m *Manager) AuthenticateJWT(tokenString string) (*Claims, error) {
	// jwt.Parseでトークンをパースし、検証用関数を渡す
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 署名方式がHMACかどうかをチェック
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		// 検証用のシークレットを返す
		return m.secret, nil
	})
}

// 任意のclaimsとsecretでJWTを発行する本番用関数
func (m *Manager) IssueJWT(claims Claims) (string, error) {
	// 1. 指定したクレームで新しいJWTトークンを作成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 2. シークレットでトークンに署名し、JWT文字列を生成
	tokenString, err := token.SignedString(m.secret)
	if err != nil {
		// 3. 署名に失敗した場合はエラーを返す
		return "", err
	}

	// 4. 正常に生成できた場合はJWT文字列を返す
	return tokenString, nil
}

// func CreateClaims(id string, duration time.Duration, issuer string) jwt.StandardClaims {
// 	return jwt.StandardClaims{
// 		Id:        id,
// 		Issuer:    issuer,
// 		ExpiresAt: time.Now().Add(duration).Unix(),
// 		IssuedAt:  time.Now().Unix(),
// 	}
// }

// func IssueJWT(claims jwt.Claims, secret string) (string, error) {
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenString, err := token.SignedString([]byte(secret))
// 	if err != nil {
// 		return "", err
// 	}
// 	return tokenString, nil
// }

// func Verify(j string, secret string) (*jwt.StandardClaims, error) {
// 	token, err := jwt.ParseWithClaims(j, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return "", errors.New("UNEXPECTED SIGNING METHOD")
// 		}
// 		return []byte(secret), nil
// 	})
// 	if err != nil {
// 		return nil, exception.ErrInvalidJWT
// 	}
// 	if err = token.Claims.Valid(); err != nil {
// 		return nil, exception.ErrInvalidJWT
// 	}
// 	standardC, ok := token.Claims.(*jwt.StandardClaims)
// 	if !ok {
// 		return nil, exception.ErrInvalidJWT
// 	}
// 	return standardC, nil
// }
