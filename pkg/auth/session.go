package auth

type Session struct {
	RefreshToken string `json:"refreshToken" bson:"refreshToken"`
	User         User
}
