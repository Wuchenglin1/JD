package dao

import (
	"JD/model"
	"database/sql"
)

const user = "root:Wcl2021214174..@tcp(110.42.165.192:3306)/RedRock_WinterWork"

var dB *sql.DB

func InitMySql() {
	db, err := sql.Open("mysql", user)
	if err != nil {
		panic(err)
	}
	dB = db
}

func SearchUserByPhone(phone string) (model.User, error) {
	u := model.User{}
	u.Phone = phone
	err := dB.QueryRow("select id,name,email,money from User where phone = ?", phone).Scan(&u.Id, &u.UserName, &u.Email, &u.Money)
	return u, err
}

func SavePhoneVerifyCode(phone, code string) error {
	_, err := dB.Exec("insert into RegisterUser values(phone=?,phoneVerifyCode=?,emailVerifyCode=?,email)", phone, code, "", "")
	if err == nil {
		return nil
	}
	_, err = dB.Exec("update RegisterUser set phoneVerifyCode = ? where phone = ?", code, phone)
	return err
}

func SearchUserByEmail(email string) (model.RegisterUser, error) {
	u := model.RegisterUser{}
	err := dB.QueryRow("select id,name,phone,email,money from User where email = ?", email).Scan(&u.Id, &u.UserName, &u.Phone, &u.Money)
	return u, err
}

func SaveEmailVerifyCode(email, code string) error {
	_, err := dB.Exec("insert into RegisterUser values(emailVerifyCode=?,emailVerifyCode,phone,phoneVerifyCode)", email, code, "", "")
	if err == nil {
		return nil
	}
	_, err = dB.Exec("update RegisterUser set emailVerifyCode = ? where email = ?", code, email)
	return err
}

func CheckVerifyCodeByEmail(u model.RegisterUser) (bool, error) {
	iU := model.RegisterUser{}
	err := dB.QueryRow("select emailVerifyCode from RegisterUser where email = ?", u.Email).Scan(&iU.VerifyCode)
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
	err := dB.QueryRow("select phoneVerifyCode from RegisterUser where phone = ?", u.Phone).Scan(&iU.VerifyCode)
	if err != nil {
		return false, err
	}
	if u.VerifyCode != iU.VerifyCode {
		return false, nil
	}
	return true, nil
}
