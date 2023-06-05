package models

type Register struct {
	Email     string `json:"email" validate:"required,email"`
	Username  string `json:"username" validate:"required,alphanum,min=3,max=15"`
	Password  string `json:"password" validate:"required,min=6,max=20"`
	FirstName string `json:"firstName" validate:"required,alphanum,min=3,max=15"`
	LastName  string `json:"lastName" validate:"required,alphanum,min=3,max=15"`
}

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=20"`
}

type UserExist struct {
	Username string `json:"username" validate:"required,alphanum,min=3,max=15"`
}

type RefreshToken struct {
	RefreshToken string `json:"refreshToken" validate:"required,jwt"`
}

type CreateProfile struct {
	Username          string   `json:"username" validate:"required,alphanum,min=3,max=15"`
	Bio               string   `json:"bio" validate:"required,max=200"`
	Language          []string `json:"language" validate:"required,min=1,max=3"`
	Category          []string `json:"category" validate:"required,min=1,max=10"`
	ApplyVerification string   `json:"applyVerification" validate:"required,boolean"`
}

type FollowProfile struct {
	UserId string `json:"userId" validate:"required,len=24"`
}

type UnfollowProfile struct {
	UserId string `json:"userId" validate:"required,len=24"`
}
