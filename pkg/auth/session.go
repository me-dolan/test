package auth

type Session struct {
	RefreshToken string
	User         User
}
