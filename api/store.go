package api

import (
	"JD/service"
	"JD/tool"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetGoods(c *gin.Context) {
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

	way := c.PostForm("way")
	switch way {
	case "0":
		str = "select gid, price, cover, name , commentamount from goods where ownerUid = ? order by  FavorableRating desc  "
	case "1":
		str = "select gid, price, cover, name , commentamount from goods where ownerUid = ? order by  FavorableRating asc  "
	case "2":
		str = "select gid, price, cover, name , commentamount from goods where ownerUid = ? order by  saleAccount desc  "
	case "3":
		str = "select gid, price, cover, name , commentamount from goods where ownerUid = ? order by  saleAccount asc  "
	case "4":
		str = "select gid, price, cover, name , commentamount from goods where ownerUid = ? order by  commentAccount desc  "
	case "5":
		str = "select gid, price, cover, name , commentamount from goods where ownerUid = ? order by  commentAccount asc  "
	case "6":
		str = "select gid, price, cover, name , commentamount from goods where ownerUid = ? order by  saleTime desc"
	case "7":
		str = "select gid, price, cover, name , commentamount from goods where ownerUid = ? order by  saleTime asc"
	case "8":
		str = "select gid, price, cover, name , commentamount from goods where ownerUid = ? order by  price desc"
	case "9":
		str = "select gid, price, cover, name , commentamount from goods where ownerUid = ? order by  price desc"
	default:
		tool.RespErrWithData(c, false, "您筛选的方式有误")
		return
	}
	m, err := service.GetGoods(str, claim.User.Id)
	if err != nil {
		if err.Error()[4:] == " no rows in result set" {
			tool.RespErrWithData(c, false, "没有查询到相关商品")
			return
		}
		fmt.Println(err)
		tool.RespErrWithData(c, false, "服务器错误")
		return
	}
	tool.RespSuccessWithData(c, m)
}

func UpdateAnnouncement(c *gin.Context) {
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

	announcement := c.PostForm("announcement")
	if announcement == "" {
		tool.RespErrWithData(c, false, "公告不能为空")
		return
	}

	if len(announcement) >= 45 {
		tool.RespErrWithData(c, false, "公告太长了")
		return
	}

	err = service.UpdateAnnouncement(claim.User.Id, announcement)
	if err != nil {
		fmt.Println(err)
		tool.RespErrWithData(c, false, "服务器错误")
		return
	}
	tool.RespSuccess(c)
}

func GetAnnouncement(c *gin.Context) {
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
	var announcement string
	announcement, err = service.GetAnnouncement(claim.User.Id)
	if err != nil {
		if err.Error()[4:] == " no rows in result set" {
			tool.RespErrWithData(c, false, "还没有设置公告，赶快添加吧")
			return
		}
		fmt.Println(err)
		tool.RespErrWithData(c, false, "服务器错误")
		return
	}
	tool.RespSuccessWithData(c, announcement)
}
