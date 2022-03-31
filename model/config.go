package model

type Config struct {
	MySql  MySql        `json:"mysql"`
	Email  Email        `json:"email"`
	Jwt    Jwt          `json:"jwt"`
	Sms    Sms          `json:"sms"`
	Oss    Oss          `json:"oss"`
	Github GitHubConfig `json:"github"`
}

type MySql struct {
	User string `json:"user"`
	Gorm string `json:"gorm"`
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

type Oss struct {
	SecretId        string `json:"secretId"`
	SecretKey       string `json:"secretKey"`
	Bucket          string `json:"bucket"`
	EndPoint        string `json:"endPoint"`
	CoverDir        string `json:"coverDir"`
	DescribeDir     string `json:"describeDir"`
	VideoDir        string `json:"videoDir"`
	DetailDir       string `json:"detailDir"`
	ColorDir        string `json:"colorDir"`
	CommentPhotoDir string `json:"commentPhotoDir"`
	CommentVideoDir string `json:"commentVideoDir"`
}
