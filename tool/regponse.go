package tool

import "github.com/gin-gonic/gin"

func RespErrWithData(c *gin.Context, status bool, data interface{}) {
	c.JSON(200, gin.H{
		"status": status,
		"data":   data,
	})
}

func RespSuccessWithData(c *gin.Context, status bool, data interface{}) {
	c.JSON(200, gin.H{
		"status": status,
		"data":   data,
	})
}

func RespSuccess(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": true,
		"data":   "",
	})
}
