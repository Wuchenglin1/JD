package dao

import "JD/model"

func GetGoods(str string, uid int) (map[int]model.GoodsInfo, error) {
	m := make(map[int]model.GoodsInfo)
	g := model.GoodsInfo{}
	err := dB.QueryRow("select name from User where uid = ?", uid).Scan(&g.OwnerName)
	if err != nil {
		return m, err
	}
	stmt, err := dB.Prepare(str)
	if err != nil {
		return m, err
	}
	defer stmt.Close()
	row, err := stmt.Query(uid)
	if err != nil {
		return m, err
	}
	defer row.Close()
	for i := 0; row.Next(); i++ {
		err = row.Scan(&g.GId, &g.Price, &g.Cover, &g.Name, &g.CommentAccount)
		if err != nil {
			return m, err
		}
		m[i] = g
	}
	return m, nil
}
