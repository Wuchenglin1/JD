package api

import (
	"JD/model"
	"JD/service"
	"JD/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"strconv"
	"strings"
	"time"
)

func Create(c *gin.Context) {
	ossCfg := tool.GetConfig().Oss
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

	g.Type, err = strconv.Atoi(_type)
	if err != nil {
		fmt.Println(err)
		tool.RespErrWithData(c, false, "服务器错误")
		return
	}
	g.Name = c.PostForm("name")
	if g.Name == "" {
		tool.RespErrWithData(c, false, "商品名称不能为空")
		return
	}
	if len(g.Name) >= 30 {
		tool.RespErrWithData(c, false, "商品名称太长啦！")
		return
	}

	price, err := strconv.Atoi(c.PostForm("price"))
	if err != nil || price <= 0 {
		fmt.Println(err)
		tool.RespErrWithData(c, false, "价格填写不正确")
		return
	}

	g.Price = price

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

	//封面入oss
	suffix := tool.GetFileSuffix(header.Filename)
	url := ossCfg.CoverDir + strconv.FormatInt(g.GId, 10) + "." + suffix
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
		url = ossCfg.DescribeDir + strconv.FormatInt(g.GId, 10) + strconv.Itoa(i) + "." + suffix
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
		url = ossCfg.VideoDir + strconv.FormatInt(g.GId, 10) + strconv.Itoa(i) + "." + suffix
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
		url = ossCfg.DetailDir + strconv.FormatInt(g.GId, 10) + strconv.Itoa(i) + "." + suffix
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

	c.JSON(200, gin.H{
		"status": true,
		"data":   "",
		"gid":    g.GId,
	})
}

// Blouse 插入一条blouse的介绍
func Blouse(c *gin.Context) {
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

	bl := model.Blouse{
		Brand:          c.PostForm("brand"),
		WomenClothing:  c.PostForm("womenClothing"),
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
	if len(bl.Brand) >= 30 || len(bl.WomenClothing) >= 30 || len(bl.Version) >= 30 || len(bl.Length) >= 30 || len(bl.SleeveLength) >= 30 || len(bl.GetModel) >= 30 || len(bl.Style) >= 30 || len(bl.Material) >= 30 || len(bl.Pattern) >= 30 || len(bl.WearingWay) >= 30 || len(bl.PopularElement) >= 30 || len(bl.SleeveType) >= 30 || len(bl.ClothesPlacket) >= 30 || len(bl.MarketTime) >= 30 || len(bl.Fabric) >= 30 || len(bl.Other) >= 30 {
		tool.RespErrWithData(c, false, "属性名称太长啦！")
		return
	}

	sa, err := strconv.Atoi(c.PostForm("suitableAge"))
	if err != nil {
		fmt.Println(err)
		tool.RespErrWithData(c, false, "请正确填写适用年龄！")
		return
	}
	bl.SuitableAge = sa
	gid, err := strconv.ParseInt(c.PostForm("gid"), 10, 64)
	if err != nil {
		tool.RespErrWithData(c, false, "")
	}
	err = service.InsertBlouse(bl, gid)
	if err != nil {
		tool.RespErrWithData(c, false, "商品错误！")
		return
	}
	tool.RespSuccess(c)
}

// ColorPhoto 为商品插入颜色和颜色的图片
func ColorPhoto(c *gin.Context) {
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

	//获取gid
	fGid := c.PostForm("gid")
	//获取颜色描述
	color := c.PostForm("color")
	if len(color) >= 15 {
		tool.RespErrWithData(c, false, "描述太长")
		return
	}
	//获取colorPhoto
	colorPhoto, header, err := c.Request.FormFile("colorPhoto")
	if err != nil {
		fmt.Println(err)
		tool.RespErrWithData(c, false, "photo上传失败")
		return
	}
	if header.Size == 0 {
		tool.RespErrWithData(c, false, "图片不能为空")
		return
	}
	if header.Size >= (5 >> 20) {
		tool.RespErrWithData(c, false, "图片太大")
		return
	}
	//存储colorPhoto
	ossCfg := tool.GetConfig().Oss
	suffix := tool.GetFileSuffix(header.Filename)
	url := ossCfg.ColorDir + fGid + "." + suffix
	err = service.SaveFile(url, colorPhoto)
	if err != nil {
		tool.RespErrWithData(c, false, "服务器错误")
		return
	}
	//存储url
	gid, err := strconv.ParseInt(fGid, 10, 64)
	err = service.InsertColorPhoto(color, url, gid)
	if err != nil {
		tool.RespErrWithData(c, false, "服务器错误")
		return
	}
	tool.RespSuccess(c)
}

func Size(c *gin.Context) {

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

	gid, err := strconv.ParseInt(c.PostForm("gid"), 10, 64)
	if err != nil {
		fmt.Println(err)
		tool.RespErrWithData(c, false, "服务器错误")
		return
	}
	size := c.PostForm("size")
	m := strings.Split(size, ";")
	err = service.InsertSize(gid, m)
	if err != nil {
		fmt.Println(err)
		tool.RespErrWithData(c, false, "服务器错误")
		return
	}
	tool.RespSuccess(c)
}

func BrowseGoods(c *gin.Context) {
	var i, str string
	i = c.PostForm("arrangement")
	switch i {
	case "0":
		str = "select gId,name,ownerUid,commentAmount,cover,price from goods order by FavorableRating desc  "
	case "1":
		str = "select gId,name,ownerUid,commentAmount,cover,price from goods order by FavorableRating asc  "
	case "2":
		str = "select gId,name,ownerUid,commentAmount,cover,price from goods order by saleAccount desc  "
	case "3":
		str = "select gId,name,ownerUid,commentAmount,cover,price from goods order by saleAccount asc  "
	case "4":
		str = "select gId,name,ownerUid,commentAmount,cover,price from goods order by commentAccount desc  "
	case "5":
		str = "select gId,name,ownerUid,commentAmount,cover,price from goods order by commentAccount asc  "
	case "6":
		str = "select gId,name,ownerUid,commentAmount,cover,price from goods order by saleTime desc"
	case "7":
		str = "select gId,name,ownerUid,commentAmount,cover,price from goods order by saleTime asc"
	case "8":
		str = "select gId,name,ownerUid,commentAmount,cover,price from goods order by price desc"
	case "9":
		str = "select gId,name,ownerUid,commentAmount,cover,price from goods order by price desc"
	default:
		tool.RespErrWithData(c, false, "您筛选的方式有误")
		return
	}
	m, err := service.BrowseGoods(str)
	if err != nil {
		tool.RespErrWithData(c, false, "服务器错误")
		return
	}
	for _, v := range m {
		tool.RespSuccessWithData(c, v)
	}
}

func GetGoodsBaseInfo(c *gin.Context) {
	gid, err := strconv.ParseInt(c.PostForm("gid"), 10, 64)
	if err != nil {
		fmt.Println(err)
		tool.RespErrWithData(c, false, "服务器错误")
		return
	}
	g, err := service.GetGoodsBaseInfo(gid)
	if err != nil {
		if err.Error()[4:] != " no rows in result set" {
			tool.RespErrWithData(c, false, "物品不存在")
			return
		}
		fmt.Println(err)
		tool.RespErrWithData(c, false, "服务器错误")
		return
	}
	tool.RespSuccessWithData(c, g)
}

func GetGoodsSize(c *gin.Context) {
	gid, err := strconv.ParseInt(c.PostForm("gid"), 10, 64)
	if err != nil {
		fmt.Println(err)
		tool.RespErrWithData(c, false, "gid填写不正确！")
		return
	}
	m, err := service.GetGoodsSize(gid)
	if err != nil {
		if err.Error()[4:] != " no rows in result set" {
			tool.RespErrWithData(c, false, "物品不存在")
			return
		}
		fmt.Println(err)
		tool.RespErrWithData(c, false, "服务器错误")
		return
	}
	for _, v := range m {
		c.JSON(200, gin.H{
			"size": v,
		})
	}
}

func GetGoodsColor(c *gin.Context) {
	gid, err := strconv.ParseInt(c.PostForm("gid"), 10, 64)
	if err != nil {
		fmt.Println(err)
		tool.RespErrWithData(c, false, "gid填写不正确！")
		return
	}
	m, err := service.GetGoodsColor(gid)
	if err != nil {
		if err.Error()[4:] != " no rows in result set" {
			tool.RespErrWithData(c, false, "物品不存在")
			return
		}
		fmt.Println(err)
		tool.RespErrWithData(c, false, "服务器错误")
		return
	}
	for _, v := range m {
		c.JSON(200, gin.H{
			"color": v.Color,
			"url":   v.Url,
		})
	}
}

func BrowseGoodsType(c *gin.Context) {

	type_, err := strconv.Atoi(c.PostForm("type"))
	if err != nil {
		fmt.Println(err)
		tool.RespErrWithData(c, false, "类型错误")
		return
	}
	m, err := service.BrowseGoodsType(type_)
	if err != nil {
		if err.Error()[4:] == " no rows in result set" {
			tool.RespErrWithData(c, false, "类型不存在！")
			return
		}
		tool.RespErrWithData(c, false, "服务器错误")
		return
	}

	for _, v := range m {
		c.JSON(200, v)
	}
}

func AddShoppingCart(c *gin.Context) {
	s := model.ShoppingCart{}
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
	//uid
	s.UId = claim.User.Id

	//获取加入商品的数量
	s.Account, err = strconv.Atoi(c.PostForm("account"))
	if err != nil {
		tool.RespErrWithData(c, false, "商品数量有误")
		return
	}
	gid, err := strconv.ParseInt(c.PostForm("gid"), 10, 64)
	if err != nil {
		tool.RespErrWithData(c, false, "商品id有误")
		return
	}
	s.Gid = gid
	s.Color = c.PostForm("color")
	s.Size = c.PostForm("size")
	s.Style = c.PostForm("style")
	if s.Color == "" {
		s.Color = "0"
	}
	if s.Size == "" {
		s.Size = "0"
	}
	if s.Style == "" {
		s.Style = "0"
	}
	//赋值价格,名称
	g, err := service.GetGoodsBaseInfo(gid)
	if err != nil {
		tool.RespErrWithData(c, false, "商品不存在！")
		return
	}
	s.Price = g.Price
	s.GoodsName = g.Name
	err = service.AddGoods(s)
	if err != nil {
		fmt.Println(err)
		tool.RespErrWithData(c, false, "服务器错误")
		return
	}
	tool.RespSuccess(c)
}

func BrowseGoodsByKeyWords(c *gin.Context) {

	keyWords := c.PostForm("keyWords")

	m, err := service.BrowseGoodsByKeyWords(keyWords)
	if err != nil {
		if err.Error()[4:] == " no rows in result set" {
			tool.RespErrWithData(c, false, "没有找到相关结果")
			return
		}
	}
	for _, v := range m {
		tool.RespSuccessWithData(c, v)
	}
}

func GoodsFocus(c *gin.Context) {
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

	f := model.GoodsFocus{}

	f.UId = claim.User.Id

	f.GId, err = strconv.Atoi(c.PostForm("gid"))
	if err != nil {
		tool.RespErrWithData(c, false, "gid不正确")
		return
	}
	f.FocusTime = time.Now()
	flag, err = service.InsertFocus(f)

	if err != nil {
		fmt.Println("插入商品错误", err)
		tool.RespErrWithData(c, false, "服务器错误")
		return
	}
	if flag == false {
		tool.RespErrWithData(c, false, "您已经关注过该商品啦！请勿重复关注")
		return
	}
	tool.RespSuccessWithData(c, "关注成功！")
}

func GetGoodsFocus(c *gin.Context) {
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

	f := model.GoodsFocus{
		UId: claim.User.Id,
	}

	m, flag, err := service.GetGoodsFocus(f)
	if err != nil {
		fmt.Println(err)
		tool.RespErrWithData(c, false, "服务器错误")
		return
	}
	if !flag {
		tool.RespErrWithData(c, false, "您还没有关注商品喔~=")
		return
	}
	for _, v := range m {
		tool.RespSuccessWithData(c, v)
	}
}
