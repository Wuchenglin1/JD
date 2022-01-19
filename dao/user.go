package dao

import (
	"JD/model"
)

func SearchUserByPhone(phone string) (model.User, error) {
	u := model.User{}
	err := dB.QueryRow("select id,name,email,money from User where phone = ?", phone).Scan(&u.Id, &u.UserName, &u.Email, &u.Money)
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

func SearchUserByEmail(email string) (model.RegisterUser, error) {
	u := model.RegisterUser{}
	err := dB.QueryRow("select id,name,phone,email,money from User where email = ?", email).Scan(&u.Id, &u.UserName, &u.Phone, &u.Email, &u.Money)
	return u, err
}

func SearchUserByUserName(userName string) (model.User, error) {
	u := model.User{}
	err := dB.QueryRow("select id,name,phone,email,money from User where name = ?", userName).Scan(&u.Id, &u.UserName, &u.Phone, &u.Email, &u.Money)
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
	_, err := dB.Exec("insert into User(name, phone, email) values (?,?,?)", u.UserName, u.Phone, u.Email)
	return err
}
