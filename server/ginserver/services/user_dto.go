package services

/*
User DTOs

DTO means Data Transfer Object
DTO is same as the API request body json and mainly used for binding it.
*/

type SignInDto struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignUpDto struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
