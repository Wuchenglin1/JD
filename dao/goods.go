package dao

import (
	"JD/model"
	"fmt"
	"time"
)

func InsertBlouse(bl model.Blouse, fGid int64) error {
	stmt, err := dB.Prepare("insert into Blouse(gid, brand, womenClothing, version, length, sleeveLength, suitableAge, getModel, style, material, pattern, wearingWay, popularElement, sleeveType, clothesPlacket, marketTime, fabric, other, time) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(fGid, bl.Brand, bl.WomenClothing, bl.Version, bl.Length, bl.SleeveLength, bl.SuitableAge, bl.GetModel, bl.Style, bl.Material, bl.Pattern, bl.WearingWay, bl.PopularElement, bl.SleeveType, bl.ClothesPlacket, bl.MarketTime, bl.Fabric, bl.Other, bl.NowTime)

	return err
}

func InsertGoods(g model.Goods, u model.User) (model.Goods, error) {
	stmt, err := dB.Prepare("insert into goods(type,name,ownerUid,price,saleTime) values (?,?,?,?,?)")
	defer stmt.Close()
	result, err := stmt.Exec(g.Type, g.Name, u.Id, g.Price, time.Now())

	if err != nil {
		fmt.Println("stmt.Exec error:", err)
	}

	g.GId, err = result.LastInsertId()
	return g, err
}

func InsertCover(g model.Goods, url string) error {
	stmt, err := dB.Prepare("update goods set cover = ? where gId = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(url, g.GId)

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func InsertDescribe(g model.Goods, url string) error {
	stmt, err := dB.Prepare("insert into photo values(gid,time,url)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(g.GId, time.Now(), url)

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func InsertVideo(g model.Goods, url string) error {
	stmt, err := dB.Prepare("insert into video values(gid,time,url)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(g.GId, time.Now(), url)

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func InsertDetail(g model.Goods, url string) error {
	stmt, err := dB.Prepare("insert into detail values(gid,time,url)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(g.GId, time.Now(), url)

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func BrowseGoods(str string) (map[int]model.GoodsInfo, error) {

	m := make(map[int]model.GoodsInfo)
	g := model.Goods{}

	stmt, err := dB.Prepare(str)
	if err != nil {
		fmt.Println(err)
		return m, err
	}
	defer stmt.Close()
	row, err := stmt.Query()
	if err != nil {
		fmt.Println(err)
		return m, err
	}
	defer row.Close()
	for i := 0; row.Next(); {
		v := m[i]
		err = row.Scan(&v.GId, &v.Name, &g.OwnerUid, &v.CommentAccount, &v.Cover, &v.Price)
		if err != nil {
			fmt.Println(err)
		}
		err = dB.QueryRow("select name from User where uid = ?", g.OwnerUid).Scan(&v.OwnerName)
		if err != nil {
			fmt.Println(err)
		}
	}
	return m, err
}

func InsertColorPhoto(color, url string, gid int64) error {
	stmt, err := dB.Prepare("insert into color(gid, color, url) values (?,?,?)")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(gid, color, url)
	if err != nil {
		fmt.Println("插入颜色错误：", err)
		return err
	}
	return nil
}

func InsertSize(gid int64, m []string) error {
	stmt, err := dB.Prepare("insert into size(gid, size) values(?,?) ")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, v := range m {
		_, err = stmt.Exec(gid, v)
		if err != nil {
			fmt.Println(err)
		}
	}
	return err
}

func GetGoodsBaseInfo(gid int64) (model.Goods, error) {
	g := model.Goods{}
	stmt, err := dB.Prepare("select gid, type, price, cover, name, owneruid, saletime, volume, commentamount, favorablerating from goods where gId = ?")
	if err != nil {
		return g, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(gid).Scan(&g.GId, &g.Type, &g.Price, &g.Cover, &g.Name, &g.OwnerUid, &g.SaleTime, &g.Volume, &g.CommentAccount, &g.FavorableRating)

	if err != nil {
		return g, err
	}
	return g, nil
}

func GetGoodsSize(gid int64) (map[int]string, error) {
	m := make(map[int]string)
	i := 0
	stmt, err := dB.Prepare("select size from size where gid = ?")
	if err != nil {
		return m, err
	}
	defer stmt.Close()
	row, err := stmt.Query(gid)
	if err != nil {
		return m, err
	}
	defer row.Close()
	for row.Next() {
		var size string
		err = row.Scan(&size)
		m[i] = size
		i++
		if err != nil {
			fmt.Println(err)
		}
	}
	return m, err
}

func GetGoodsColor(gid int64) (map[int]model.GoodsColor, error) {
	m := make(map[int]model.GoodsColor)
	i := 0
	stmt, err := dB.Prepare("select color,url from color where gid = ?")
	if err != nil {
		return m, err
	}
	defer stmt.Close()
	row, err := stmt.Query(gid)
	if err != nil {
		return m, err
	}
	for row.Next() {
		var col model.GoodsColor
		err = row.Scan(&col.Color, &col.Url)
		m[i] = col
		if err != nil {
			fmt.Println(err)
		}
	}
	return m, err
}

func BrowseGoodsType(type_ int) (map[int]model.GoodsInfo, error) {
	m := make(map[int]model.GoodsInfo)
	stmt, err := dB.Prepare("select gId,cover,price,name,commentAmount,ownerUid from goods where type = ?")
	var g model.GoodsInfo
	var ownerUid int
	if err != nil {
		fmt.Println("BrowseGoodsType Error :", err)
		return m, err
	}
	defer stmt.Close()
	result, err := stmt.Query(type_)

	if err != nil {
		fmt.Println("BrowseGoodsType Error :", err)
		return m, err
	}
	defer result.Close()
	for i := 0; result.Next(); i++ {
		err = result.Scan(&g.GId, &g.Cover, &g.Price, &g.Name, &g.CommentAccount, &ownerUid)
		if err != nil {
			fmt.Println(err)
		}
		//查询商家名字
		err = dB.QueryRow("select name from User where uid = ?", ownerUid).Scan(&g.OwnerName)
		m[i] = g
		if err != nil {
			fmt.Println(err)
		}
	}
	return m, err
}

func AddGoods(s model.ShoppingCart) error {
	stmt, err := dB.Prepare("select  goodsName from shoppingCart where uId = ? and gid = ?")
	defer stmt.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = stmt.QueryRow(s.UId, s.Gid).Scan(&s.GoodsName)
	if err != nil {
		if err.Error()[4:] == " no rows in result set" {
			//商品不存在
			stmt1, err1 := dB.Prepare("insert  into shoppingCart(uid, gid, goodsname, color, size, style, price, account) values (?,?,?,?,?,?,?,?)")
			defer stmt1.Close()
			if err1 != nil {
				fmt.Println(err1)
				return err1
			}
			_, err = stmt1.Exec(s.UId, s.Gid, s.GoodsName, s.Color, s.Size, s.Style, s.Price, s.Account)
			return err
		}
		return err
	}
	//商品存在,数量+1
	stmt, err = dB.Prepare("update shoppingCart set account=account+1 where gid = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_ = stmt.QueryRow(s.Gid)
	return nil
}

func BrowseGoodsByKeyWords(keyWords string) (map[int]model.GoodsInfo, error) {
	m := make(map[int]model.GoodsInfo)
	g := model.GoodsInfo{}
	//模糊查询
	stmt, err := dB.Prepare("select gid, price, cover, name, owneruid, commentamount from goods where name like  ?")
	if err != nil {
		return m, err
	}
	defer stmt.Close()
	//这里需要符合like的用法
	row, err := stmt.Query("%" + keyWords + "%")
	if err != nil {
		return m, err
	}
	defer row.Close()
	for i := 0; row.Next(); i++ {
		var uid int
		row.Scan(&g.GId, &g.Price, &g.Cover, &g.Name, &uid, &g.CommentAccount)
		//查询商家名称
		err = dB.QueryRow("select  name from User where uid = ?", uid).Scan(&g.OwnerName)
		if err != nil {
			fmt.Println(err)
		}
		m[i] = g
	}
	return m, err
}
