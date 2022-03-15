package service

import (
	"JD/dao"
	"JD/model"
	"JD/tool"
	"golang.org/x/crypto/bcrypt"
)

func Authorization(userInfo model.UserInfo) (model.User, bool, error) {
	u := model.User{}
	err := dao.CheckUserByGithubAccount(userInfo)
	if err != nil && err.Error()[4:] != " no rows in result set" {
		return u, false, err
	}
	//如果没有绑定过github,注册账号
	if err.Error()[4:] == " no rows in result set" {

		password, err1 := tool.CreateRandomString(15)
		if err1 != nil {
			return u, false, err1
		}
		name, err2 := tool.CreateRandomString(15)
		if err2 != nil {
			return u, false, err2
		}
		u = model.User{
			GithubAccount: userInfo.Id,
			Password:      password,
			Name:          name,
			Avatar:        userInfo.AvatarUrl,
		}
		hashPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		u.Password = string(hashPassword)
		err2 = dao.CreateOauthUser(u)
		return u, false, err
	}
	return u, true, nil
}
