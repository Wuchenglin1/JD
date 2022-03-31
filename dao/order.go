package dao

import (
	"JD/model"
	"gorm.io/gorm"
)

func CreateOrder(o model.Order) (bool, error) {

	err := db.Transaction(func(tx *gorm.DB) error {
		resDB := tx.Create(&o)
		if resDB.Error != nil {
			return resDB.Error
		}
		for _, v := range o.Settlement {
			g := model.Goods{}
			resDB = db.Where("id = ?", v.GId).First(&g)
			if resDB.Error != nil {
				return resDB.Error
			}
			if g.Inventory <= v.Account {
				return nil
			}
			g.Inventory = g.Inventory - v.Account
			resDB = db.Save(&g)
			if resDB.Error != nil {
				return resDB.Error
			}
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func CheckAllOrder(o model.Order) ([]model.Order, error) {
	var arr []model.Order
	resDB := db.Preload("settlements").Where("uid = ?", o.Uid).Find(&arr)
	if resDB.Error != nil {
		return nil, resDB.Error
	}

	return arr, nil
}

func CheckSpecified(o model.Order) (model.Order, error) {
	resDB := db.Preload("settlements").Where("orderNumber = ?", o.OrderNumber).First(&o)
	if resDB.Error != nil {
		return o, resDB.Error
	}
	return o, nil
}

func CancelOrder(o model.Order) error {
	db.First(&o)
	err := db.Transaction(func(tx *gorm.DB) error {
		resDB := tx.Model(&o).Select("status").Updates(map[string]interface{}{"status": "已取消"})
		if resDB.Error != nil {
			return resDB.Error
		}
		for k := range o.Settlement {
			g := model.Goods{}
			resDB = db.Where("id = ?", o.Settlement[k].GId).First(&g)
			if resDB.Error != nil {
				return resDB.Error
			}
			g.Inventory += o.Settlement[k].Account
			resDB = db.Save(&g)
			if resDB.Error != nil {
				return resDB.Error
			}
		}
		return nil
	})
	return err
}

func SolveOrder(o model.Order, u model.User) (bool, error) {
	err := db.Transaction(func(tx *gorm.DB) error {
		resDB := tx.Where("id = ?", u.ID).First(&u)
		if resDB.Error != nil {
			return resDB.Error
		}
		u.Money = u.Money - o.TotalPrice
		resDB = db.Save(&u)
		if resDB.Error != nil {
			return resDB.Error
		}
		for _, v := range o.Settlement {
			g := model.Goods{}
			resDB = db.Where("id = ?", v.GId).First(&g)
			if resDB.Error != nil {
				return resDB.Error
			}
			resDB = db.Where("id = ?", g.OwnerUid).First(&u)
			if resDB.Error != nil {
				return resDB.Error
			}
			u.Money += v.Account * v.Price
			resDB = db.Save(&u)
			if resDB.Error != nil {
				return resDB.Error
			}
			g.Volume += v.Account
			resDB = db.Save(&g)
			if resDB.Error != nil {
				return resDB.Error
			}
		}
		resDB = db.First(&o)
		if resDB.Error != nil {
			return resDB.Error
		}
		if o.Status == "已超时" || o.Status == "已取消" {
			return gorm.ErrRecordNotFound
		}
		o.Status = "待收货"
		resDB = db.Save(&o)
		if resDB.Error != nil {
			return resDB.Error
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	return true, nil

}

func CreateConsigneeInfo(c model.ConsigneeInfo) error {
	resDB := db.Create(&c)
	return resDB.Error
}

func GetConsigneeInfo(c model.ConsigneeInfo) ([]model.ConsigneeInfo, error) {
	var arr []model.ConsigneeInfo
	resDB := db.Where("uid = ?", c.Uid).Find(&arr)
	if resDB.Error != nil {
		return nil, resDB.Error
	}
	return arr, resDB.Error
}

func DeleteConsigneeInfo(c model.ConsigneeInfo) error {
	resDB := db.Where("id = ?", c.ID).First(&c)
	if resDB.Error != nil {
		return resDB.Error
	}
	return nil
}

func ConfirmOrder(order model.Order) error {
	resDB := db.Where("orderNumber = ?", order.OrderNumber).First(&order)
	if resDB.Error != nil {
		return resDB.Error
	}
	return nil
}

func DeleteOrder(order model.Order) error {
	resDB := db.Where("orderNumber = ?", order.OrderNumber).First(&order)
	if resDB.Error != nil {
		return resDB.Error
	}
	return nil
}

func CheckOrderByStatus(o model.Order, status string) ([]model.Order, error) {
	var arr []model.Order
	resDB := db.Where("status = ?", o.Status).Find(&arr)
	if resDB.Error != nil {
		return nil, resDB.Error
	}

	return arr, nil
}
