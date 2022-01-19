package tool

import "github.com/gin-gonic/gin"

func RespErrWithData(c *gin.Context, status bool, data string) {
	c.JSON(200, gin.H{
		"status": status,
		"date":   data,
	})
}
