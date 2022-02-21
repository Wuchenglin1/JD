package dao

import (
	"JD/tool"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var dB *sql.DB

func InitMySql() {
	config := tool.GetConfig()
	db, err := sql.Open("mysql", config.MySql.User)
	if err != nil {
		panic(err)
	}
	dB = db
}

func UpdateAnnouncement(uid int, announcement string) error {
	var a string
	err := dB.QueryRow("select announcement from store where uid = ?", uid).Scan(&a)
	if err != nil {
		if err.Error()[4:] == " no rows in result set" {
			stmt, err := dB.Prepare("insert into store (uid, announcement) values (?,?)")
			if err != nil {
				return err
			}
			_, err = stmt.Exec(uid, announcement)
			if err != nil {
				return err
			}
		}
		return err
	}
	stmt, err := dB.Prepare("update store set announcement = ? where uid = ?")
	if err != nil {
		return err

	}
	defer stmt.Close()
	_, err = stmt.Exec(announcement, uid)
	if err != nil {
		return err
	}
	return nil

}

func GetAnnouncement(uid int) (string, error) {
	var announcement string
	stmt, err := dB.Prepare("select announcement from store where uid = ? ")
	if err != nil {
		return "", err
	}
	defer stmt.Close()
	err = stmt.QueryRow(uid).Scan(&announcement)
	if err != nil {
		return "", err
	}
	return announcement, nil
}
