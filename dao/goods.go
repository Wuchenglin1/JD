package dao

import (
	"JD/model"
	"fmt"
	"time"
)

func InsertBlouse(bl model.Blouse, g model.Goods) (int64, error) {
	stmt, err := dB.Prepare("insert into Blouse(fGid, price, brand, womenClothing, size, color, version, length, sleeveLength, suitableAge, getModel, style, material, pattern, wearingWay, popularElement, sleeveType, clothesPlacket, marketTime, fabric, other, time) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(g.GId, bl.Price, bl.Brand, bl.WomenClothing, bl.Size, bl.Color, bl.Version, bl.Length, bl.SleeveLength, bl.SuitableAge, bl.GetModel, bl.Style, bl.Material, bl.Pattern, bl.WearingWay, bl.PopularElement, bl.SleeveType, bl.ClothesPlacket, bl.MarketTime, bl.Fabric, bl.Other, bl.NowTime)

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func InsertGoods(g model.Goods, u model.User) (model.Goods, error) {
	stmt, err := dB.Prepare("insert into goods(type,name,ownerUid) values (?,?,?)")
	defer stmt.Close()
	result, err := stmt.Exec(g.Type, g.Name, u.Id)

	if err != nil {
		fmt.Println("stmt.Exec error:", err)
	}

	g.GId, err = result.LastInsertId()
	return g, err
}

func InsertCover(g model.Goods, url string) error {
	stmt, err := dB.Prepare("insert into cover values(type,id,time,url)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(g.Type, g.GId, time.Now(), url)

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func InsertDescribe(g model.Goods, url string) error {
	stmt, err := dB.Prepare("insert into photo values(type,id,time,url)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(g.Type, g.GId, time.Now(), url)

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func InsertVideo(g model.Goods, url string) error {
	stmt, err := dB.Prepare("insert into video values(type,id,time,url)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(g.Type, g.GId, time.Now(), url)

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func InsertDetail(g model.Goods, url string) error {
	stmt, err := dB.Prepare("insert into detail values(type,id,time,url)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(g.Type, g.GId, time.Now(), url)

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
