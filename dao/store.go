package dao

import (
	"JD/model"
	"gorm.io/gorm"
)

func GetGoods(str string, uid int) ([]model.Goods, error) {
	var arr []model.Goods
	g := model.Goods{}
	u := model.User{}
	resDB := db.Where("id = ?", uid).First(&u)
	if resDB.Error != nil {
		return arr, resDB.Error
	}
	g.OwnerName = u.UserName
	switch str {
	case "0":
		db.Order("FavorableRating desc,id").First(&arr)
	case "1":
		db.Order("FavorableRating asc").First(&arr)
	case "2":
		db.Order("saleAccount desc").First(&arr)
	case "3":
		db.Order("saleAccount asc").First(&arr)
	case "4":
		db.Order("commentAccount desc").First(&arr)
	case "5":
		db.Order("commentAccount asc").First(&arr)
	case "6":
		db.Order("saleTime desc").First(&arr)
	case "7":
		db.Order("saleTime asc").First(&arr)
	case "8":
		db.Order("price desc").First(&arr)
	case "9":
		db.Order("price asc").First(&arr)
	default:
		return arr, gorm.ErrRecordNotFound
	}
	return arr, nil
}
