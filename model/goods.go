package model

import "time"

// GoodsInfo 专门用来返回给前端一个商品基本信息
type GoodsInfo struct {
	GId            int    `json:"GId"`
	Cover          string `json:"cover"`
	Price          int    `json:"price"`
	Name           string `json:"name"`
	CommentAccount int    `json:"commentAccount"`
	OwnerName      string `json:"ownerName"`
}

type Goods struct {
	Type            int    `json:"type"`
	Name            string `json:"name"`
	GId             int64  `json:"gid"`
	Price           int    `json:"price"`
	OwnerUid        int    `json:"ownerUid"`
	OwnerName       string `json:"ownerName"`
	CommentAccount  int    `json:"commentAccount"`
	Volume          int    `json:"volume"`
	FavorableRating int    `json:"FavorableRating"`
	Cover           string `json:"cover"`
}

//Blouse 女士衬衫
type Blouse struct {
	Gid            int       `json:"gid"`
	Brand          string    `json:"brand"`
	WomenClothing  string    `json:"WomenClothing"`
	Version        string    `json:"version"`
	Length         string    `json:"length"`
	SleeveLength   string    `json:"sleeveLength"`
	SuitableAge    int       `json:"suitableAge"`
	GetModel       string    `json:"getModel"`
	Style          string    `json:"style"`
	Material       string    `json:"material"`
	Pattern        string    `json:"pattern"`
	WearingWay     string    `json:"wearingWay"`
	PopularElement string    `json:"popularElement"`
	SleeveType     string    `json:"sleeveType"`
	ClothesPlacket string    `json:"clothesPlacket"`
	MarketTime     string    `json:"marketTime"`
	Fabric         string    `json:"fabric"`
	Other          string    `json:"other"`
	NowTime        time.Time `json:"now_time"`
}

type CowboyPants struct {
	Gid            int       `json:"gid"`
	Brand          string    `json:"brand"`
	WaistType      string    `json:"waistType"`
	Height         string    `json:"height"`
	Pants          string    `json:"pants"`
	Thick          int       `json:"thick"`
	Stretch        string    `json:"stretch"`
	Material       string    `json:"material"`
	SuitableAge    int       `json:"suitableAge"`
	MarketTime     string    `json:"marketTime"`
	PopularElement string    `json:"popularElement"`
	Fabric         string    `json:"fabric"`
	FrontPants     string    `json:"frontPants"`
	NowTime        time.Time `json:"nowTime"`
}

type Color struct {
}
