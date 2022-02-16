package api

import (
	"JD/model"
	"JD/service"
	"JD/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"strconv"
	"time"
)

func AddComment(c *gin.Context) {
	//解析token
	claim, err := service.ParseAccessToken(c.PostForm("token"))
	if err != nil {
		if err.Error() == "token contains an invalid number of segments" {
			c.JSON(200, "token错误！")
			return
		}
	}
	flag, str := service.CheckTokenErr(claim, err)
	if !flag {
		tool.RespErrWithData(c, false, str)
		return
	}

	var header *multipart.FileHeader
	comment := model.Comment{
		UId:  claim.User.Id,
		Time: time.Now(),
	}
	//读取商品的gid
	gid := c.PostForm("gid")
	comment.GId, err = strconv.Atoi(gid)
	if err != nil {
		tool.RespErrWithData(c, false, "gid有误")
		return
	}
	comment.Comment = c.PostForm("comment")

	//读取是否匿名
	isAnonymous := c.PostForm("isAnonymous")
	switch isAnonymous {
	case "1":
		comment.IsAnonymous = true
	case "2":
		comment.IsAnonymous = false
	default:
		tool.RespErrWithData(c, false, "是否匿名")
		return
	}

	//读取评分
	comment.Grade, err = strconv.Atoi(c.PostForm("grade"))
	if err != nil {
		tool.RespErrWithData(c, false, "评分有误")
		return
	}
	//读取视频
	comment.Video, header, err = c.Request.FormFile("video")
	if err != nil {
		tool.RespErrWithData(c, false, "上传的视频有误")
		return
	}
	if header.Size >= 50<<20 {
		tool.RespErrWithData(c, false, "视频文件太大")
		return
	}
	//读取评论的图片到map中
	form := c.Request.MultipartForm

	files := form.File["photo[]"]

	for _, file := range files {
		if file.Size >= 10<<20 {
			tool.RespErrWithData(c, false, "图片太大啦")
			return
		}
		f, err1 := file.Open()
		if err1 != nil {
			fmt.Println(err1)
			tool.RespErrWithData(c, false, err1)
			return
		}
		comment.Photo[file.Filename] = f
	}
	err = service.SaveComment(comment, header.Filename)
	if err != nil {
		fmt.Println(err)
		tool.RespErrWithData(c, false, err)
		return
	}
	tool.RespSuccess(c)
}

func ViewComment(c *gin.Context) {
	gid, err := strconv.Atoi(c.PostForm("gid"))
	if err != nil {
		tool.RespErrWithData(c, false, "gid错误")
		return
	}
	comment := model.Comment{
		GId: gid,
	}
	m, err := service.ViewComment(comment)
	if err != nil {
		if err.Error()[4:] == " no rows in result set" {
			tool.RespErrWithData(c, false, "没有查询到评论")
			return
		}
		fmt.Println(err)
		tool.RespErrWithData(c, false, "服务器错误")
		return
	}
	for _, v := range m {
		tool.RespSuccessWithData(c, v)
	}
}

func ReplyComment(c *gin.Context) {
	//解析token
	claim, err := service.ParseAccessToken(c.PostForm("token"))
	if err != nil {
		if err.Error() == "token contains an invalid number of segments" {
			c.JSON(200, "token错误！")
			return
		}
	}
	flag, str := service.CheckTokenErr(claim, err)
	if !flag {
		tool.RespErrWithData(c, false, str)
		return
	}

	commentId := c.PostForm("commentId")
	comment := model.Comment{UId: claim.User.Id}
	comment.CommentId, err = strconv.Atoi(commentId)
	if err != nil {
		tool.RespErrWithData(c, false, "评论id有误")
		return
	}
	commentStr := c.PostForm("comment")
	if commentStr == "" {
		tool.RespErrWithData(c, false, "评论不能为空")
		return
	}
	comment.Comment = commentStr

	comment.GId, err = strconv.Atoi(c.PostForm("gid"))
	if err != nil {
		tool.RespErrWithData(c, false, "gid有误")
	}

	err = service.ReplyComment(comment)
	if err != nil {
		if err.Error()[4:] == " no rows in result set" {
			tool.RespErrWithData(c, false, "没有找到该评论")
			return
		}
		fmt.Println(err)
		tool.RespErrWithData(c, false, "服务器错误")
		return
	}
	tool.RespSuccess(c)
}

func ViewSpecificComment(c *gin.Context) {
	commentId, err := strconv.Atoi(c.PostForm("commentId"))
	if err != nil {
		tool.RespErrWithData(c, false, "commentId有误")
		return
	}
	comment := model.Comment{CommentId: commentId}
	comment, err = service.ViewSpecificComment(comment)
	if err != nil {
		if err.Error()[4:] == " no rows in result set" {
			tool.RespErrWithData(c, false, "没有多余的评论啦")
			return
		}
		tool.RespErrWithData(c, false, "服务器错误")
		return
	}
	tool.RespSuccessWithData(c, comment)
}
