package token

import "github.com/me-dolan/test/internal/user"

type Session struct {
	RefreshToken string `json:"refreshToken" bson:"refreshToken"`
	User         user.User
}
