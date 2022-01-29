package service

import (
	"JD/dao"
	"JD/model"
)

func InsertBlouse(bl model.Blouse, fGid int64) error {
	return dao.InsertBlouse(bl, fGid)
}

func InsertGoods(g model.Goods, u model.User) (model.Goods, error) {
	return dao.InsertGoods(g, u)
}

func InsertCover(g model.Goods, url string) error {
	return dao.InsertCover(g, url)
}

func InsertDescribe(g model.Goods, url string) error {
	return dao.InsertDescribe(g, url)
}

func InsertVideo(g model.Goods, url string) error {
	return dao.InsertVideo(g, url)
}

func InsertDetail(g model.Goods, url string) error {
	return dao.InsertDetail(g, url)
}

func BrowseGoods(str string) (map[int]model.GoodsInfo, error) {
	return dao.BrowseGoods(str)
}

func InsertColorPhoto(color, url string, gid int64) error {
	return dao.InsertColorPhoto(color, url, gid)
}

func InsertSize(gid int64, m []string) error {
	return dao.InsertSize(gid, m)
}

func GetGoodsBaseInfo(gid int64) (model.Goods, error) {
	return dao.GetGoodsBaseInfo(gid)
}

func GetGoodsSize(gid int64) (map[int]string, error) {
	return dao.GetGoodsSize(gid)
}

func GetGoodsColor(gid int64) (map[int]model.GoodsColor, error) {
	return dao.GetGoodsColor(gid)
}

func BrowseGoodsType(type_ int) (map[int]model.GoodsInfo, error) {
	return dao.BrowseGoodsType(type_)
}

func AddGoods(s model.ShoppingCart) error {
	return dao.AddGoods(s)
}

func BrowseGoodsByKeyWords(keyWords string) (map[int]model.GoodsInfo, error) {
	//查询所有商品
	if keyWords == "" {
		str := "select gId,name,ownerUid,commentAmount,cover,price from goods order by saleTime desc"
		return dao.BrowseGoods(str)
	}
	//查询关键词商品
	return dao.BrowseGoodsByKeyWords(keyWords)
}
