package token

import (
	"encoding/base64"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/me-dolan/test/internal/user"
	"go.mongodb.org/mongo-driver/mongo"
)

type Tokens struct {
	db        *mongo.Client
	SecretKey string
}

type AuthTokens struct {
	AccesToken   string
	RefreshToken string
}

type AuthTokenClaim struct {
	*jwt.StandardClaims
	User user.User
}

func (t *Tokens) generateTokens(guid string) (AuthTokens, user.User, error) {
	var u user.User
	token := jwt.NewWithClaims(jwt.SigningMethodES512, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
	})

	// token := jwt.New(jwt.SigningMethodHS512) // SHA512
	// expiresAt := time.Now().Add(time.Minute * 1).Unix()
	// token.Claims = &AuthTokenClaim{
	// 	&jwt.StandardClaims{
	// 		ExpiresAt: expiresAt,
	// 	},
	// 	u,
	// }

	tokenString, err := token.SignedString([]byte(t.SecretKey))
	if err != nil {
		return AuthTokens{}, u, err
	}

	refreshToken := uuid.New()
	refreshTokenBase64 := base64.StdEncoding.EncodeToString([]byte(refreshToken.String()))

	return AuthTokens{
		AccesToken:   tokenString,
		RefreshToken: refreshTokenBase64,
	}, u, nil
}
