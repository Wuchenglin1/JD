package dao

import "JD/model"

func CheckUserByGithubAccount(userInfo model.UserInfo) error {
	u := model.UserInfo{}
	resDB := db.Where("githubAccount = ?", userInfo.Id).First(&u)
	if resDB.Error != nil {
		return resDB.Error
	}
	return nil

}

func CreateOauthUser(u model.User) error {
	resDB := db.Create(&u)
	if resDB.Error != nil {
		return resDB.Error
	}
	return nil
}
