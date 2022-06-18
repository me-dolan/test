package domain

type User struct {
	Guid         string `json:"guid" bson:"guid"`
	RefreshToken string `json:"refreshtoken" bson:"refreshtoken"`
}
