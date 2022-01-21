package model

type Config struct {
	MySql MySql `json:"mysql"`
	Email Email `json:"email"`
	Jwt   Jwt   `json:"jwt"`
	Sms   Sms   `json:"sms"`
}

type MySql struct {
	User string `json:"user"`
}

type Email struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
}

type Sms struct {
	SecretId   string `json:"secretId"`
	SecretKey  string `json:"secretKey"`
	AppId      string `json:"appId"`
	AppKey     string `json:"appKey"`
	SignId     string `json:"signId"`
	TemplateId string `json:"templateId"`
	Sign       string `json:"sign"`
}
