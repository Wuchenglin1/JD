package dao

import (
	"JD/model"
)

func SearchUserByPhone(phone string) (model.User, error) {
	u := model.User{}
	err := dB.QueryRow("select uid,name,email,money,password from User where phone = ?", phone).Scan(&u.Id, &u.UserName, &u.Email, &u.Money, &u.Password)
	u.Phone = phone
	return u, err
}

func SavePhoneVerifyCode(phone, code string) error {
	_, err := dB.Exec("insert into phoneVerify(phone, verifyCode) values(?,?)", phone, code)
	if err == nil {
		return nil
	}
	_, err = dB.Exec("update phoneVerify set verifyCode = ? where phone = ?", code, phone)
	return err
}

func SearchUserByEmail(email string) (model.User, error) {
	u := model.User{}
	err := dB.QueryRow("select uid,name,phone,email,money,password from User where email = ?", email).Scan(&u.Id, &u.UserName, &u.Phone, &u.Email, &u.Money, &u.Password)
	return u, err
}

func SearchUserByUserName(userName string) (model.User, error) {
	u := model.User{}
	err := dB.QueryRow("select uid,name,phone,email,money,password from User where name = ?", userName).Scan(&u.Id, &u.UserName, &u.Phone, &u.Email, &u.Money, &u.Password)
	return u, err
}

func SaveEmailVerifyCode(email, code string) error {
	_, err := dB.Exec("insert into emailVerify(email,verifyCode) values(?,?)", email, code)
	if err == nil {
		return nil
	}
	_, err = dB.Exec("update emailVerify set verifyCode = ? where email = ?", code, email)
	return err
}

func CheckVerifyCodeByEmail(u model.RegisterUser) (bool, error) {
	iU := model.RegisterUser{}
	err := dB.QueryRow("select verifyCode from emailVerify where email = ?", u.Email).Scan(&iU.VerifyCode)
	if err != nil {
		return false, err
	}
	if u.VerifyCode != iU.VerifyCode {
		return false, nil
	}
	return true, nil
}

func CheckVerifyCodeByPhone(u model.RegisterUser) (bool, error) {
	iU := model.RegisterUser{}
	err := dB.QueryRow("select verifyCode from phoneVerify where phone = ?", u.Phone).Scan(&iU.VerifyCode)
	if err != nil {
		return false, err
	}
	if u.VerifyCode != iU.VerifyCode {
		return false, nil
	}
	return true, nil
}

func SaveUser(u model.RegisterUser) error {
	_, err := dB.Exec("insert into User(name, phone, email,password) values (?,?,?,?)", u.UserName, u.Phone, u.Email, u.Password)
	return err
}

func BrowseShoppingCart(uid int) (map[int]model.ShoppingCart, int, error) {
	var totalPrice = 0
	i := 0
	m := make(map[int]model.ShoppingCart)
	stmt, err := dB.Prepare("select uid, gid, goodsname, color, size, price, account,cover from shoppingCart where uId = ?")
	if err != nil {
		return m, 0, err
	}
	defer stmt.Close()
	row, err := stmt.Query(uid)
	if err != nil {
		return m, 0, err
	}
	defer row.Close()
	for row.Next() {
		g := model.ShoppingCart{}
		err = row.Scan(&g.UId, &g.Gid, &g.GoodsName, &g.Color, &g.Size, &g.Price, &g.Account, &g.Cover)
		if err != nil {
			return m, 0, err
		}
		totalPrice += g.Price * g.Account
		m[i] = g
		i++
	}

	return m, totalPrice, nil
}

func SearchUserByUid(u model.User) (model.User, error) {
	stmt, err := dB.Prepare("select uid, name, password, headPic, phone, email, money from User where uid = ?")
	if err != nil {
		return u, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(u.Id).Scan(&u.Id, &u.UserName, &u.Password, u.HeadPic, &u.Phone, &u.Email, &u.Money)
	return u, err
}

func UpdateUserByUid(u model.User) error {
	stmt, err := dB.Prepare("update User set uid=?,name=?,password=?,headPic=?,phone=?,email=?,money=?  where uid = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(u.Id, u.UserName, u.Password, u.HeadPic, u.Phone, u.Email, u.Money)
	return err
}
