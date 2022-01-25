package model

import "time"

type Goods struct {
	Type            string `json:"type"`
	Name            string `json:"name"`
	GId             int64  `json:"gid"`
	OwnerUid        int    `json:"ownerUid"`
	CommentAccount  int    `json:"commentAccount"`
	Volume          int    `json:"volume"`
	FavorableRating int    `json:"FavorableRating"`
}

//Blouse 女士衬衫
type Blouse struct {
	Gfid           int       `json:"gfid"`
	Gid            int       `json:"gid"`
	Price          int       `json:"price"`
	Brand          string    `json:"brand"`
	WomenClothing  string    `json:"WomenClothing"`
	Size           string    `json:"size"`
	Color          string    `json:"color"`
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
	Gfid           int       `json:"gfid"`
	Gid            int       `json:"gid"`
	Price          int       `json:"price"`
	Brand          string    `json:"brand"`
	Size           string    `json:"size"`
	Color          string    `json:"color"`
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
