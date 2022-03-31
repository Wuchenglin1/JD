package dao

import (
	"JD/model"
	"gorm.io/gorm"
)

func InsertBlouse(bl model.Blouse, fGid int64) error {
	bl.ID = uint(fGid)
	resDB := db.Create(&bl)
	return resDB.Error
}

func InsertGoods(g model.Goods, u model.User) (model.Goods, error) {
	g.OwnerUid = u.ID
	resDB := db.Create(g)
	if resDB.Error != nil {
		return g, resDB.Error
	}
	resDB = db.Last(&g)
	return g, resDB.Error
}

func InsertCover(g model.Goods, url string) error {
	resDB := db.Where("id = ?", g.ID).First(&g)
	if resDB.Error != nil {
		return resDB.Error
	}
	g.Cover = url
	resDB = db.Save(g)
	return resDB.Error
}

func InsertDescribe(g model.Goods, url string) error {
	resDB := db.Create(&g)
	return resDB.Error
}

func InsertVideo(g model.Goods, url string) error {
	g.VideoUrl = url
	resDB := db.Create(&g)
	return resDB.Error
}

func InsertDetail(g model.Goods, url string) error {
	goods := model.GoodsDetail{}
	goods.ID = g.ID
	resDB := db.Create(&goods)
	return resDB.Error
}

func BrowseGoods(str string) ([]model.Goods, error) {
	var arr []model.Goods
	resDB := db.Find(&arr)
	if resDB.Error != nil {
		return arr, resDB.Error
	}
	for k := range arr {
		u := model.User{}
		resDB = db.Where("id = ?", arr[k].OwnerUid).First(&u)
		if resDB.Error != nil {
			return arr, resDB.Error
		}
		arr[k].OwnerName = u.UserName
	}
	return arr, nil
}

func InsertColorPhoto(color, url string, gid int64) error {
	g := model.GoodsColor{}
	g.ID = uint(gid)
	g.Url = url
	g.Color = color
	resDB := db.Save(&g)
	return resDB.Error
}

func InsertSize(gid int64, m []string) error {
	s := model.Size{}
	s.ID = uint(gid)
	s.Size = m
	resDB := db.Create(&s)
	return resDB.Error
}

func GetGoodsBaseInfo(gid int64) (model.Goods, error) {
	g := model.Goods{}
	resDB := db.Where("id = ?", gid).First(&g)
	if resDB.Error != nil {
		return g, resDB.Error
	}
	return g, nil

}

func GetGoodsSize(gid int64) ([]string, error) {
	s := model.Size{}
	resDB := db.Where("id = ?", gid).First(&s)
	if resDB.Error != nil {
		return nil, resDB.Error
	}
	return s.Size, nil

}

func GetGoodsColor(gid int64) ([]model.GoodsColor, error) {
	var arr []model.GoodsColor
	resDB := db.Where("id = ?", gid).Find(&arr)
	return arr, resDB.Error
}

func BrowseGoodsType(type_ int) ([]model.Goods, error) {
	var arr []model.Goods
	resDB := db.Where("type = ?").Find(&arr)
	if resDB.Error != nil {
		return nil, resDB.Error
	}
	for k := range arr {
		u := model.User{}
		resDB = db.Where("id = ?", arr[k].OwnerUid).First(u)
		if resDB.Error != nil {
			return nil, resDB.Error
		}
		arr[k].OwnerName = u.UserName
	}
	return arr, nil
}

func AddGoods(s model.ShoppingCart) error {
	s1 := model.ShoppingCart{}
	resDB := db.Where("gid = ? and id = ?", s.Gid, s.ID).First(&s1)
	if resDB.Error != nil {
		if resDB.Error == gorm.ErrRecordNotFound {
			resDB = db.Create(&s)
			if resDB.Error != nil {
				return resDB.Error
			}
		}
		return resDB.Error
	}
	s1.Account += s.Account
	resDB = db.Save(&s1)
	return resDB.Error
}

func BrowseGoodsByKeyWords(keyWords string) ([]model.Goods, error) {
	var arr []model.Goods
	resDB := db.Where("name like ?", keyWords).Find(&arr)
	if resDB.Error != nil {
		return nil, resDB.Error
	}
	for k := range arr {
		g := model.Goods{}
		resDB = db.Where("id = ?", arr[k].ID).First(&g)
		if resDB.Error != nil {
			return nil, resDB.Error
		}
	}
	return arr, nil
}

func InsertFocus(f model.GoodsFocus) (bool, error) {
	resDB := db.Where("id = ?", f.ID).First(&f)
	if resDB.Error != nil {
		if resDB.Error == gorm.ErrRecordNotFound {
			resDB = db.Create(&f)
			if resDB.Error != nil {
				return false, resDB.Error
			}
			return true, nil
		}
		return false, resDB.Error
	}
	return false, nil
}

func GetGoodsFocus(f model.GoodsFocus) ([]model.GoodsFocus, bool, error) {
	var arr []model.GoodsFocus
	resDB := db.Order("desc time").Where("id = ?", f.ID).Find(&arr)
	if resDB.Error != nil {
		return nil, false, resDB.Error
	}
	return arr, true, nil
}

func DeleteFocus(f model.GoodsFocus) (bool, error) {
	resDB := db.Where("id = ? and gid = ?", f.ID, f.GoodsBaseInfo.ID).First(&f)
	if resDB.Error != nil {
		if resDB.Error == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, resDB.Error
	}
	resDB = db.Delete(&f)
	if resDB.Error != nil {
		return false, resDB.Error
	}
	return true, nil
}

func DeleteShoppingCart(s model.ShoppingCart) (bool, error) {
	resDB := db.Where("gid = ? and id = ?", s.Gid, s.ID).First(&s)
	if resDB.Error != nil {
		if resDB.Error == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, resDB.Error
	}
	resDB = db.Delete(&s)
	if resDB.Error != nil {
		return false, resDB.Error
	}
	return true, nil
}

func CreateGoods(goods model.Goods, describePhoto string, detailPhoto string) (model.Goods, error) {
	goods.DescribePhotoUrl = describePhoto
	goods.DetailPhotoUrl = detailPhoto
	err := db.Transaction(func(tx *gorm.DB) error {
		resDB := tx.Create(&goods)
		if resDB.Error != nil {
			return resDB.Error
		}
		return nil
	})
	return goods, err
}
