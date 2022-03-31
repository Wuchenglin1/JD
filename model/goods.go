package model

import (
	"gorm.io/gorm"
	"time"
)

type GoodsBaseInfo struct {
	gorm.Model
	Name  string
	Price float64
}

// GoodsInfo 专门用来返回给前端一个商品基本信息
//type GoodsInfo struct {
//	GId            int    `json:"gId"`
//	Cover          string `json:"cover"`
//	Price          int    `json:"price"`
//	Name           string `json:"name"`
//	CommentAccount int    `json:"commentAccount"`
//	OwnerName      string `json:"ownerName"`
//}

type Goods struct {
	GoodsBaseInfo
	Type             int       `json:"type"`
	Inventory        int       `json:"inventory"`
	OwnerUid         uint      `json:"ownerUid"`
	OwnerName        string    `json:"ownerName"`
	SaleTime         time.Time `json:"saleTime"`
	CommentAccount   int       `json:"commentAccount"`
	Volume           int       `json:"volume"`
	FavorableRating  int       `json:"FavorableRating"`
	DetailPhotoUrl   string
	DescribePhotoUrl string
	Cover            string `json:"cover"`
	VideoUrl         string
}

type GoodsDetail struct {
	GoodsBaseInfo
	Detail string
}

//Blouse 女士衬衫
type Blouse struct {
	GoodsBaseInfo
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
	NowTime        time.Time `json:"nowTime"`
}

type CowboyPants struct {
	GoodsBaseInfo
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

type GoodsColor struct {
	GoodsBaseInfo
	Color string `json:"color,"`
	Url   string `json:"url""`
}

type Size struct {
	GoodsBaseInfo
	Size []string
}

type GoodsFocus struct {
	GoodsBaseInfo
	UId             int       `json:"uId"`
	Cover           string    `json:"cover"`
	CommentAccount  int       `json:"commentAccount"`
	FavorableRating int       `json:"favorableRating"`
	FocusTime       time.Time `json:"focusTime"`
}
