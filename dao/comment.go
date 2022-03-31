package dao

import (
	"JD/model"
	"gorm.io/gorm"
)

func SaveComment(c model.Comment) (int64, error) {

	g := model.Goods{}
	err := db.Transaction(func(tx *gorm.DB) error {
		//先插入评论
		resTX := tx.Create(&c)
		if resTX.Error != nil {
			return resTX.Error
		}
		//评论数+1
		tx.Where("id = ?", c.GId).First(&g)
		g.CommentAccount++
		tx.Save(&g)
		return nil
	})
	if err != nil {
		return 0, err
	}
	return int64(g.ID), nil

}

func SaveCommentPhoto(url string, i int64) error {
	resDB := db.Model(&model.Comment{}).Select("photoUrl", "id").Updates(map[string]interface{}{"url": url, "id": i})
	return resDB.Error
}

func SaveCommentVideo(url string, i int64) error {
	resDB := db.Model(&model.Comment{}).Select("videoUrl", "id").Updates(map[string]interface{}{"url": url, "id": i})
	return resDB.Error
}

func ViewComment(c model.Comment) (map[int]model.Comment, error) {
	m := make(map[int]model.Comment)
	var arr []model.Comment
	db.Where("gid = ? and fCommentId = ?", c.GId, "-1").Find(arr)
	for k := range arr {
		if arr[k].IsAnonymous {
			arr[k].Name = "匿名用户"
		}
		var arr1 []model.Comment
		c = model.Comment{}
		db.Where("fCommentId = ?", arr[k].ID).First(&arr1)
		if db.Error != nil {
			if db.Error == gorm.ErrRecordNotFound {
				return m, nil
			}
			return m, db.Error
		}
		m[k] = arr[k]
	}
	return m, nil
}

func ReplyComment(c model.Comment) error {
	db.Transaction(func(tx *gorm.DB) error {
		tx.Create(&c)
		g := model.Goods{}
		resDB := tx.Where("commentId = ?", c.ID).First(&g)
		if resDB.Error != nil {
			return resDB.Error
		}
		g.CommentAccount++
		tx.Save(&g)
		return nil
	})
	return nil
}

var num = 0

func ViewSpecificComment(c model.Comment) (map[int]model.Comment, error) {
	m := make(map[int]model.Comment)
	var arr []model.Comment
	db.Where("fCommentId = ?", c.ID).First(&arr)
	for k := range arr {
		err := ViewSonComment(m, arr[k])
		if err != nil {
			return m, err
		}
	}
	return m, nil
}

func ViewSonComment(m map[int]model.Comment, c model.Comment) error {
	var arr []model.Comment
	resDB := db.Where("fCommentId = ?", c.ID).First(&arr)
	if resDB.Error != nil {
		if resDB.Error == gorm.ErrRecordNotFound {
			return nil
		}
		return resDB.Error
	}
	for k := range arr {
		m[num] = arr[k]
		err := ViewSonComment(m, arr[k])
		if err != nil {
			return err
		}
		num++
	}
	return nil
}
