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

func Blouse(c *gin.Context) {
	ossCfg := tool.GetConfig().Oss
	//解析token
	claim, err := service.ParseAccessToken(c.PostForm("token"))
	flag, str := service.CheckTokenErr(claim, err)
	if !flag {
		tool.RespErrWithData(c, false, str)
		return
	}
	u := claim.User

	//商品部分
	g := model.Goods{}
	_type := c.PostForm("type")
	if _type != "0520101" {
		tool.RespErrWithData(c, false, "类型不正确")
		return
	}
	//channalSlice := []string{}
	//flag = false
	//for _, channelType := range channalSlice {
	//	if channelType == _type {
	//		flag = true
	//		break
	//	}
	//}
	//if flag == false {
	//	tool.RespErrWithData(c, false, "类型不正确")
	//	return
	//}

	g.Type = _type

	g.Name = c.PostForm("name")
	if g.Name == "" {
		tool.RespErrWithData(c, false, "商品名称不能为空")
		return
	}
	if len(g.Name) >= 30 {
		tool.RespErrWithData(c, false, "商品名称太长啦！")
		return
	}
	bl := model.Blouse{
		Brand:          c.PostForm("brand"),
		WomenClothing:  c.PostForm("womenClothing"),
		Size:           c.PostForm("size"),
		Color:          c.PostForm("color"),
		Version:        c.PostForm("version"),
		Length:         c.PostForm("length"),
		SleeveLength:   c.PostForm("sleeveLength"),
		GetModel:       c.PostForm("getModel"),
		Style:          c.PostForm("style"),
		Material:       c.PostForm("Material"),
		Pattern:        c.PostForm("pattern"),
		WearingWay:     c.PostForm("wearingWay"),
		PopularElement: c.PostForm("popularElement"),
		SleeveType:     c.PostForm("sleeveType"),
		ClothesPlacket: c.PostForm("clothesPlacket"),
		MarketTime:     c.PostForm("marketTime"),
		Fabric:         c.PostForm("fabric"),
		Other:          c.PostForm("other"),
		NowTime:        time.Now(),
	}
	if len(bl.Brand) >= 30 || len(bl.WomenClothing) >= 30 || len(bl.Size) >= 30 || len(bl.Color) >= 30 || len(bl.Version) >= 30 || len(bl.Length) >= 30 || len(bl.SleeveLength) >= 30 || len(bl.GetModel) >= 30 || len(bl.Style) >= 30 || len(bl.Material) >= 30 || len(bl.Pattern) >= 30 || len(bl.WearingWay) >= 30 || len(bl.PopularElement) >= 30 || len(bl.SleeveType) >= 30 || len(bl.ClothesPlacket) >= 30 || len(bl.MarketTime) >= 30 || len(bl.Fabric) >= 30 || len(bl.Other) >= 30 {
		tool.RespErrWithData(c, false, "属性名称太长啦！")
		return
	}
	price, err := strconv.Atoi(c.PostForm("price"))
	if err != nil || price <= 0 {
		fmt.Println(err)
		tool.RespErrWithData(c, false, "价格填写不正确")
		return
	}
	bl.Price = price
	sa, err := strconv.Atoi(c.PostForm("suitableAge"))
	if err != nil {
		fmt.Println(err)
		tool.RespErrWithData(c, false, "请正确填写适用年龄！")
		return
	}
	bl.SuitableAge = sa

	//读取商品封面
	coverFile, header, err := c.Request.FormFile("cover")
	if err != nil {
		tool.RespErrWithData(c, false, "cover上传失败")
		return
	}

	//封面为空
	if header.Size == 0 {
		tool.RespErrWithData(c, false, "封面文件不能为空的啦！")
		return
	}
	//封面限制10mb大小
	if header.Size > (10 << 20) {
		tool.RespErrWithData(c, false, "封面文件太大啦！")
		return
	}

	p := make(map[string]multipart.File)

	//读取商品展示的图片 存储到map中
	form := c.Request.MultipartForm
	files := form.File["describePhoto[]"]
	for _, file := range files {
		if file.Size == 0 {
			tool.RespErrWithData(c, false, "商品展示图不能为空呀！")
			return
		}
		if file.Size >= (30 << 20) {
			tool.RespErrWithData(c, false, "商品展示图太大啦！")
			return
		}
		f, err1 := file.Open()
		if err1 != nil {
			fmt.Println("file open error1 :", err1)
			tool.RespErrWithData(c, false, "服务器错误")
			return
		}
		p[file.Filename] = f
	}

	//读取商品展示的视频 存储到map中
	v := make(map[string]multipart.File)

	files = form.File["describeVideo[]"]
	for _, file := range files {
		if file.Size == 0 {
			tool.RespErrWithData(c, false, "商品展示视频不能为空")
			return
		}
		if file.Size >= (1 >> 30) {
			tool.RespErrWithData(c, false, "商品展示视频太大了")
			return
		}
		f, err1 := file.Open()
		if err1 != nil {
			fmt.Println("file open error2 :", err1)
			tool.RespErrWithData(c, false, "服务器错误")
			return
		}
		v[file.Filename] = f
	}

	//读取商品介绍图片 存储到map中
	d := make(map[string]multipart.File)

	files = form.File["detailPhoto[]"]
	for _, file := range files {
		if file.Size == 0 {
			tool.RespErrWithData(c, false, "商品介绍不能为空")
			return
		}
		if file.Size >= (10 >> 20) {
			tool.RespErrWithData(c, false, "商品介绍图片太大")
		}
		f, err1 := file.Open()
		if err1 != nil {
			fmt.Println("file open error3 :", err1)
			tool.RespErrWithData(c, false, "服务器错误")
			return
		}
		d[file.Filename] = f
	}

	//货品入mysql
	g, err = service.InsertGoods(g, u)
	if err != nil {
		tool.RespErrWithData(c, false, "服务器错误")
		return
	}
	id, err := service.InsertBlouse(bl, g)

	//封面入oss
	suffix := tool.GetFileSuffix(header.Filename)
	url := ossCfg.CoverDir + strconv.FormatInt(id, 10) + "." + suffix
	err = service.SaveFile(url, coverFile)
	if err != nil {
		fmt.Println("封面入oss错误:", err)
		tool.RespErrWithData(c, false, "服务器错误")
		return
	}
	//封面url入mysql
	err = service.InsertCover(g, url)
	if err != nil {
		fmt.Println(err)
		tool.RespErrWithData(c, false, "服务器错误")
		return
	}

	//商品图片入oss
	i := 0
	for k, v := range p {
		suffix = tool.GetFileSuffix(k)
		url = ossCfg.DescribeDir + strconv.FormatInt(id, 10) + strconv.Itoa(i) + "." + suffix
		i++
		err = service.SaveFile(url, v)
		if err != nil {
			fmt.Println("商品图片入oss错误:", err)
			tool.RespErrWithData(c, false, "服务器错误")
			return
		}
		//商品图片url入mysql
		err = service.InsertDescribe(g, url)
		if err != nil {
			fmt.Println(err)
			tool.RespErrWithData(c, false, "服务器错误")
			return
		}
	}

	//视频入oss
	i = 0
	for k, v := range v {
		suffix = tool.GetFileSuffix(k)
		url = ossCfg.VideoDir + strconv.FormatInt(id, 10) + strconv.Itoa(i) + "." + suffix
		err = service.SaveFile(url, v)
		if err != nil {
			fmt.Println("商品视频插入oss错误:", err)
			tool.RespErrWithData(c, false, "服务器错误")
			return
		}
		//视频url入mysql
		err = service.InsertVideo(g, url)
		if err != nil {
			fmt.Println(err)
			tool.RespErrWithData(c, false, "服务器错误")
			return
		}
	}

	//详情图片入oss
	i = 0
	for k, v := range d {
		suffix = tool.GetFileSuffix(k)
		url = ossCfg.DetailDir + strconv.FormatInt(id, 10) + strconv.Itoa(i) + "." + suffix
		err = service.SaveFile(url, v)
		if err != nil {
			fmt.Println("详情图片插入oss错误:", err)
			tool.RespErrWithData(c, false, "服务器错误")
			return
		}
		//详情图片url入mysql
		err = service.InsertDetail(g, url)
		if err != nil {
			fmt.Println(err)
			tool.RespErrWithData(c, false, "服务器错误")
			return
		}
	}
}