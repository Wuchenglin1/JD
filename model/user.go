package model

import "gorm.io/gorm"

type UserBaseInfo struct {
	gorm.Model
	UserName string `json:"userName"`
}

type User struct {
	UserBaseInfo
	Password      string      `json:"password"`
	Phone         string      `json:"phone"`
	Email         string      `json:"email"`
	Favorite      string      `json:"favorite"`
	ShoppingCart  map[int]int `json:"shoppingCart"`
	Money         int         `json:"money"`
	HeadPic       string      `json:"headPic"`
	GithubAccount int         `json:"githubAccount"`
	Avatar        string      `json:"avatar"`
}

type ShoppingCart struct {
	UserBaseInfo
	Gid        int64   `json:"gid"`
	GoodsName  string  `json:"goodsName"`
	Cover      string  `json:"cover"`
	Color      string  `json:"color"`
	Size       string  `json:"size"`
	Price      float64 `json:"price"`
	Account    int     `json:"account"`
	TotalPrice int     `json:"totalPrice"`
}

type Announcement struct {
	UserBaseInfo
	Announcements string `json:"announcements"`
}
