package service

import (
	"JD/dao"
	"JD/model"
)

func GetGoods(str string, uid int) (map[int]model.GoodsInfo, error) {
	return dao.GetGoods(str, uid)
}
