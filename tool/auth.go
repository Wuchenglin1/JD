package tool

import (
	"JD/model"
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
)

func CreateRandomString(len int) (string, error) {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, err := rand.Int(rand.Reader, bigInt)
		if err != nil {
			return "", err
		}
		container += string(str[randomInt.Int64()])
	}
	return container, nil
}

func GetUserInfo(accessToken string) (model.UserInfo, error) {
	info := model.UserInfo{}
	client := &http.Client{}

	request, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return info, err
	}
	request.Header.Set("Authorization", fmt.Sprintf("token %s", accessToken))

	resp, err := client.Do(request)
	if err != nil {
		return info, err
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

			return info, err2
		}
		//直接绑定对象
		err = json.Unmarshal(buf[:n], &info)
		if err != nil {
			return info, err
		}
	}
	return info, nil
}
