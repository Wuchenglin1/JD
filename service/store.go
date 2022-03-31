package service

import (
	"JD/dao"
	"JD/model"
)

func GetGoods(str string, uid int) (map[int]model.Goods, error) {
	return dao.GetGoods(str, uid)
}

func UpdateAnnouncement(uid int, announcement string) error {
	return dao.UpdateAnnouncement(uid, announcement)
}

func GetAnnouncement(uid int) (string, error) {
	return dao.GetAnnouncement(uid)
}
