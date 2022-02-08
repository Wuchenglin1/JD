package model

import (
	"time"
)

type Order struct {
	Uid         int          `json:"uid"`         //创建订单的用户
	OrderNumber string       `json:"orderNumber"` //订单号
	Consignee   string       `json:"consignee"`   //收货人名字
	Address     string       `json:"address"`     //收货地址
	Phone       string       `json:"phone"`       //电话
	PayWay      string       `json:"payWay"`      //支付方式
	TotalPrice  int          `json:"totalPrice"`  //总价格
	Status      string       `json:"status"`      //订单状态
	Time        time.Time    `json:"time"`        //创建订单的时间
	Settlement  []Settlement `json:"settlement"`
}

type Settlement struct {
	GId     int64  `json:"gid"`
	Name    string `json:"name"`
	Cover   string `json:"cover"`
	Price   int    `json:"price"`
	Account int    `json:"account"`
	Color   string `json:"color"`
	Size    string `json:"size"`
}
