package dao

import (
	"JD/model"
)

func SearchUserByPhone(phone string) (model.User, error) {
	u := model.User{}
	resDB := db.Where("phone = ?", phone).First(&u)
	return u, resDB.Error
}

func SearchUserByEmail(email string) (model.User, error) {
	u := model.User{}
	resDB := db.Where("email = ?", email).First(&u)
	return u, resDB.Error
}

func SearchUserByUserName(userName string) (model.User, error) {
	u := model.User{}
	resDB := db.Where("userName = ?", userName).First(&u)
	return u, resDB.Error
}

func SaveUser(u model.User) error {
	resDB := db.Create(&u)
	return resDB.Error
}

func BrowseShoppingCart(uid int) ([]model.ShoppingCart, float64, error) {
	var totalPrice = 0.0
	var arr []model.ShoppingCart
	resDB := db.Where("id = ?", uid).Find(&arr)
	if resDB.Error != nil {
		return arr, totalPrice, resDB.Error
	}
	for k := range arr {
		totalPrice += arr[k].Price
	}
	return arr, totalPrice, nil
}

func SearchUserByUid(u model.User) (model.User, error) {
	resDB := db.Where("id = ?", u.ID).First(&u)
	return u, resDB.Error
}

func RechargeBalance(u model.User, money int) error {
	resDB := db.Where("id = ?", u.ID).First(&u)
	if resDB.Error != nil {
		return resDB.Error
	}
	u.Money += money
	resDB = db.Save(&u)
	return resDB.Error
}
