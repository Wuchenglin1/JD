package api

import (
	"JD/service"
	"JD/tool"
	"github.com/gin-gonic/gin"
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

}
