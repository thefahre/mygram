package requests

import (
	"time"

	_ "github.com/asaskevich/govalidator"
)

type UserRegisterRequest struct {
	Username string `json:"username" form:"username" valid:"required~Username is required"`
	Email    string `json:"email" form:"email" valid:"required~Email is required,email~Invalid format"`
	Password string `json:"password" form:"password" valid:"required~Password is required,minstringlength(6)~Password has to be a minimum length of six cahracters"`
	Age      uint   `json:"age" form:"age" valid:"required~Age is required"`
}

type UserRegisterResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Age      uint   `json:"age" form:"age"`
}

type UserLoginRequest struct {
	Email    string `json:"email" form:"email" valid:"required~Email is required,email~Invalid format"`
	Password string `json:"password" form:"password" valid:"required~Password is required,minstringlength(6)~Password has to be a minimum length of six cahracters"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}

type CreatePhotoRequest struct {
	Title    string `json:"title" valid:"required"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url" valid:"required"`
}

type CreateCommentRequest struct {
	Message string `json:"message"`
}

type CreateCommentResponse struct {
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}

type CreateSocialmediaRequest struct {
	Name           string `json:"name" valid:"required"`
	SocialMediaURL string `json:"social_media_url"`
}
