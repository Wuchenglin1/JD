package model

type User struct {
	Id       int    `json:"id" :"id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Money    int    `json:"money"`
}

type RegisterUser struct {
	Id         int    `json:"id" :"id"`
	UserName   string `json:"user_name"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	VerifyCode string `json:"verifyCode"`
	Money      int    `json:"money"`
}
