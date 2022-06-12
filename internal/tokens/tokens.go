package tokens

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

func NewToken(db *mongo.Client, secretKey string) *Tokens {
	return &Tokens{
		db:        db,
		SecretKey: secretKey,
	}
}

type AuthTokens struct {
	AccesToken   string `json:"acces_token"`
	RefreshToken string `json:"refresh_token"`
}

type AuthTokenClaim struct {
	*jwt.StandardClaims
	User user.User
}

func (t *Tokens) generateTokens(guid string) (AuthTokens, user.User, error) {
	var user user.User
	user.Guid = guid
	
	expiresAt := time.Now().Add(time.Minute * 1).Unix()

	token := jwt.New(jwt.SigningMethodHS512) // SHA512
	token.Claims = &AuthTokenClaim{
		&jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
		user,
	}

	tokenString, err := token.SignedString([]byte(t.SecretKey))
	if err != nil {
		return AuthTokens{}, user, err
	}

	refreshToken := uuid.New()
	refreshTokenBase64 := base64.StdEncoding.EncodeToString([]byte(refreshToken.String()))

	return AuthTokens{
		AccesToken:   tokenString,
		RefreshToken: refreshTokenBase64,
	}, user, nil
}
