package dao

import (
	"JD/model"
	"JD/tool"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB

func InitMySql() {
	config := tool.GetConfig().MySql
	dB, err := gorm.Open(mysql.Open(config.Gorm), &gorm.Config{
		SkipDefaultTransaction:                   false,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println(err)
		return
	}

	sqlDB.SetMaxIdleConns(10)           //连接池中最大的空闲连接数
	sqlDB.SetMaxOpenConns(10)           //连接池最多容纳的链接数量
	sqlDB.SetConnMaxLifetime(time.Hour) //连接池中链接的最大可复用时间

	db = dB

}

func UpdateAnnouncement(uid int, announcement string) error {
	ann := model.Announcement{}
	resDB := db.Where("id = ?", uid).First(&ann)
	if err := resDB.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ann.ID = uint(uid)
			ann.Announcements = announcement
			db.Create(&ann)
		}
		return err
	}
	ann.Announcements = announcement
	resDB = db.Save(&ann)

	return resDB.Error
}

func GetAnnouncement(uid int) (string, error) {
	ann := model.Announcement{}
	resDB := db.Where("id = ?", uid).First(ann)
	if err := resDB.Error; err != nil {
		return "", err
	}
	return ann.Announcements, nil
}
