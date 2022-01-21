package model

type Config struct {
	MySql MySql `json:"mysql"`
	Email Email `json:"email"`
	Jwt   Jwt   `json:"jwt"`
}

type MySql struct {
	User string `json:"user"`
}

type Email struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
}
