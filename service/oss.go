package service

import (
	"JD/tool"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"mime/multipart"
)

func SaveFile(url string, file multipart.File) error {
	ossCfg := tool.GetConfig().Oss
	client, err := oss.New(ossCfg.EndPoint, ossCfg.SecretId, ossCfg.SecretKey)
	if err != nil {
		return err
	}

	bucket, err := client.Bucket(ossCfg.Bucket)
	if err != nil {
		return err
	}
	err = bucket.PutObject(url, file)
	return err
}
