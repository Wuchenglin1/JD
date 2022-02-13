package service

import (
	"JD/dao"
	"JD/model"
	"JD/tool"
	"strconv"
)

func SaveComment(c model.Comment, videoName string) error {
	i, err := dao.SaveComment(c)
	if err != nil {
		return err
	}
	//获取config
	ossCfg := tool.GetConfig().Oss
	//存入图片到oss和mysql
	for k, v := range c.Photo {
		suffix := tool.GetFileSuffix(k)
		url := ossCfg.CommentPhotoDir + strconv.Itoa(c.GId) + strconv.FormatInt(i, 10) + "." + suffix
		//url入mysql
		err = dao.SaveCommentPhoto(url, i)
		if err != nil {
			return err
		}
		err = SaveFile(url, v)
		if err != nil {
			return err
		}
	}
	//存入视频到oss和mysql

	suffix := tool.GetFileSuffix(videoName)
	url := ossCfg.CommentPhotoDir + strconv.Itoa(c.GId) + strconv.FormatInt(i, 10) + "." + suffix
	//url入mysql
	err = dao.SaveCommentVideo(url, i)
	if err != nil {
		return err
	}
	err = SaveFile(url, c.Video)
	if err != nil {
		return err
	}
	return nil
}

func ViewComment(c model.Comment) (map[int]model.Comment, error) {
	return dao.ViewComment(c)
}

func ReplyComment(c model.Comment) error {
	return dao.ReplyComment(c)
}
