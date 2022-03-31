package api

import (
	"JD/model"
	"JD/service"
	"JD/tool"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func Authorize(c *gin.Context) {
	config := tool.GetConfig().Github
	state, err := tool.CreateRandomString(20)
	if err != nil {
		fmt.Println(err)
		return
	}
	config.State = state
	fmt.Println("config:", config)
	fmt.Println(state)
	c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s&state=%s", config.ClientId, config.RedirectUrl, config.State))
}

func Authorization(c *gin.Context) {
	config := tool.GetConfig().Github

	code, flag := c.GetQuery("code")
	config.Code = code
	fmt.Println(code, flag)

	//获取token
	token, err := GetAuthToken(code)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("token:", token)
	//这里进行一次打印token(为了好看)
	fmt.Println("accessToken:", token.AccessToken)
	fmt.Println("tokenType:", token.TokenType)
	fmt.Println("scope:", token.Scope)
	//通过accessToken拿取用户的信息
	userInfo, err := tool.GetUserInfo(token.AccessToken)
	if err != nil {
		fmt.Println("Get UserInfo error:", err)
		return
	}

	fmt.Println(userInfo)
	u := model.User{}
	u, flag, err = service.Authorization(userInfo)
	if err != nil && err.Error()[4:] != " no rows in result set" {
		fmt.Println(err)
		tool.RespErrWithData(c, false, "服务器错误")
		return
	}
	//如果没有绑定账号，直接返回初始账号和密码以及一个refreshToken和token
	if err.Error()[4:] == " no rows in result set" {
		tool.RespSuccessWithData(c, gin.H{
			"name":     u.Name,
			"password": u.Password,
		})
	}

	//绑定了账号，直接登录

	//创建一个存在一天的refreshToken
	rt, err := service.CreateToken(u, 86400, "refreshToken")
	if err != nil {
		fmt.Println(err)
		tool.RespErrWithData(c, false, "服务器错误")
		return
	}
	//创建一个存在5min的token
	t, err := service.CreateToken(u, 300, "token")
	if err != nil {
		fmt.Println(err)
		tool.RespErrWithData(c, false, "服务器错误")
		return
	}
	c.JSON(200, gin.H{
		"status":       true,
		"data":         "登录成功",
		"token":        t,
		"refreshToken": rt,
	})
	//c.Redirect(http.StatusMovedPermanently, "http://110.42.165.192:8080")
}

func GetAuthToken(code string) (model.TokenResponse, error) {

	config := tool.GetConfig().Github

	token := new(model.TokenResponse)

	client := &http.Client{}

	req, _ := http.NewRequest("POST", fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%v&client_secret=%v&code=%v", config.ClientId, config.ClientSecret, code), nil)

	//可以改变header来改变传入的值形式,可以尝试绑定为json对象
	req.Header.Set("accept", "application/json")
	resp, err1 := client.Do(req)

	if err1 != nil {
		return *token, err1
	}
	defer resp.Body.Close()
	buf := make([]byte, 4096)
	for {
		//读取数据
		n, err2 := resp.Body.Read(buf)
		//读完就退出循环
		if n == 0 {
			break
		}
		if err2 != nil && err2 != io.EOF {
			return *token, err2
		}
		//绑定json对象
		err := json.Unmarshal(buf[:n], token)
		if err != nil {
			return *token, err
		}
	}
	return *token, nil
}
