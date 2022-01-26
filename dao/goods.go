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
	stmt, err := dB.Prepare("insert into color(fGid, color, url) values (?,?,?)")
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

func InsertSize(gid int64, size string) error {
	stmt, err := dB.Prepare("insert into size(gid, size) values(?,?) ")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(gid, size)
	return err
}
