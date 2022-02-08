package dao

import (
	"JD/model"
	"fmt"
	"log"
)

func CreateOrder(o model.Order) error {
	tx, err := dB.Begin()
	if err != nil {
		log.Fatalln("开启事务失败：", err)
		return err
	}
	stmt, err := tx.Prepare("insert into goodsOrder(uid, orderNumber, consignee, address, phone, payWay,totalPrice, time) values (?,?,?,?,?,?,?,?)")
	if err != nil {
		err1 := tx.Rollback()
		if err1 != nil {
			fmt.Println(err1)
		}
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(o.Uid, o.OrderNumber, o.Consignee, o.Address, o.Phone, o.PayWay, o.TotalPrice, o.Time)
	if err != nil {
		if err != nil {
			err1 := tx.Rollback()
			if err1 != nil {
				fmt.Println(err1)
			}
			return err
		}
	}

	for _, v := range o.Settlement {

		stmt1, err := tx.Prepare("insert into orderGoods(orderNumber, gid, name, price,account,color, size) values(?,?,?,?,?,?,? )")
		if err != nil {
			if err != nil {
				err1 := tx.Rollback()
				if err1 != nil {
					fmt.Println(err1)
				}
				return err
			}
		}
		defer stmt1.Close()

		_, err = stmt.Exec(o.OrderNumber, v.GId, v.Name, v.Price, v.Account, v.Color, v.Size)
		if err != nil {
			if err != nil {
				err1 := tx.Rollback()
				if err1 != nil {
					fmt.Println(err1)
				}
				return err
			}
		}
	}
	err = tx.Commit()

	return err
}

func CheckAllOrder(o model.Order) (map[int]model.Order, error) {
	m := make(map[int]model.Order)
	stmt, err := dB.Prepare("select orderNumber, consignee, address, phone, payWay, totalPrice, time, status from goodsOrder where uid = ?")
	if err != nil {
		return m, err
	}
	defer stmt.Close()
	row, err := stmt.Query(o.Uid)
	if err != nil {
		return m, err
	}
	defer row.Close()
	for i := 0; row.Next(); i++ {
		err = row.Scan(&o.OrderNumber, &o.Consignee, &o.Address, &o.Phone, &o.PayWay, &o.TotalPrice, &o.Time, &o.Status)
		if err != nil {
			return m, err
		}
		if o.PayWay == "1" {
			o.PayWay = "在线支付"
		}
		length := 0
		err = dB.QueryRow("select count(*) from orderGoods where orderNumber = ?", o.OrderNumber).Scan(&length)
		if err != nil {
			return m, err
		}
		array := make([]model.Settlement, length)
		//查询每个订单下的每个商品
		rows, err1 := dB.Query("select gid, name, account, color, size from orderGoods where orderNumber = ?", o.OrderNumber)
		if err1 != nil {
			return m, err1
		}
		for k := 0; rows.Next(); k++ {
			err1 = rows.Scan(&array[k].GId, &array[k].Name, &array[k].Account, &array[k].Color, &array[k].Size)
			if err1 != nil {
				return m, err1
			}
			err1 = dB.QueryRow("select cover from goods where gid  = ?", array[k].GId).Scan(&array[k].Cover)
			if err1 != nil {
				return m, err1
			}
			o.Settlement[k] = array[k]
		}
		rows.Close()
		m[i] = o
	}
	return m, nil
}

func CheckSpecified(o model.Order) (model.Order, error) {
	stmt, err := dB.Prepare("select orderNumber, consignee, address, phone, payWay, totalPrice, time, status from goodsOrder where uid = ?")
	if err != nil {
		return o, err
	}
	defer stmt.Close()
	_ = stmt.QueryRow(o.Uid).Scan(&o.OrderNumber, &o.Consignee, &o.Address, &o.Phone, &o.PayWay, &o.TotalPrice, &o.Time, &o.Status)

	//先查询该订单下有多少商品,以确定切片的长度
	length := 0
	err = dB.QueryRow("select count(*) from orderGoods where orderNumber = ?", o.OrderNumber).Scan(&length)
	if err != nil {
		return o, err
	}
	array := make([]model.Settlement, length)
	//查询每个订单下的每个商品
	rows, err1 := dB.Query("select gid, name, account, color, size from orderGoods where orderNumber = ?", o.OrderNumber)
	if err1 != nil {
		return o, err1
	}
	for k := 0; rows.Next(); k++ {
		err1 = rows.Scan(&array[k].GId, &array[k].Name, &array[k].Account, &array[k].Color, &array[k].Size)
		if err1 != nil {
			return o, err1
		}
		err1 = dB.QueryRow("select cover from goods where gid  = ?", array[k].GId).Scan(&array[k].Cover)
		if err1 != nil {
			return o, err1
		}
		o.Settlement[k] = array[k]
	}
	return o, nil
}

func CancelOrder(o model.Order) error {
	_, err := dB.Exec("update goodsOrder set status = ? where orderNumber = ?", "已取消", o.OrderNumber)
	return err
}
