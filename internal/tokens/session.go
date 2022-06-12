package tokens

import "github.com/me-dolan/test/internal/user"

type Session struct {
	RefreshToken string `json:"refreshToken"`
	User         user.User
}
