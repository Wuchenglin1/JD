package model

type User struct {
	Id            int         `json:"id"`
	UserName      string      `json:"user_name"`
	Password      string      `json:"password"`
	Phone         string      `json:"phone"`
	Email         string      `json:"email"`
	Favorite      string      `json:"favorite"`
	ShoppingCart  map[int]int `json:"shoppingCart"`
	Money         int         `json:"money"`
	HeadPic       string      `json:"headPic"`
	GithubAccount int         `json:"githubAccount"`
	Name          string      `json:"name"`
	Avatar        string      `json:"avatar"`
}

type RegisterUser struct {
	Id         int    `json:"id"`
	UserName   string `json:"user_name"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	VerifyCode string `json:"verifyCode"`
	Money      int    `json:"money"`
}

type ShoppingCart struct {
	UId        int    `json:"uId"`
	Gid        int64  `json:"gid"`
	GoodsName  string `json:"goodsName"`
	Cover      string `json:"cover"`
	Color      string `json:"color"`
	Size       string `json:"size"`
	Price      int    `json:"price"`
	Account    int    `json:"account"`
	TotalPrice int    `json:"totalPrice"`
}
