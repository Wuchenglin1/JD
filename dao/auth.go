package dao

import "JD/model"

func CheckUserByGithubAccount(userInfo model.UserInfo) error {
	stmt, err := dB.Prepare("select uid from User where githubAccount = ? ")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(userInfo.Id)
	if err != nil {
		return err
	}
	return nil
}

func CreateOauthUser(u model.User) error {
	stmt, err := dB.Prepare("insert into User (name,password,githubAccount,avatar) values (?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(u.Name, u.Password, u.GithubAccount, u.Avatar)
	if err != nil {
		return err
	}
	return nil
}
