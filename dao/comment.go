package dao

import (
	"JD/model"
	"fmt"
	"time"
)

func SaveComment(c model.Comment) (int64, error) {
	//开启事务
	tx, err := dB.Begin()
	if err != nil {
		fmt.Println("开启事务失败:", err)
		return 0, err
	}
	//插入评论到mysql
	stmt, err := tx.Prepare("insert into goodsComment(gid,uid,comment,grade,isAnonymous,time) values(?,?,?,?,?,?)")
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			fmt.Println("回滚失败:", err)
			return 0, err
		}
		return 0, err
	}
	defer stmt.Close()
	row, err := stmt.Exec(c.GId, c.UId, c.Comment, c.Grade, c.IsAnonymous, c.Time)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			fmt.Println("回滚失败:", err)
			return 0, err
		}
		return 0, err
	}
	_, err = dB.Exec("update goods set commentAmount = commentAmount +1")
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			fmt.Println("回滚失败:", err)
			return 0, err
		}
		return 0, err
	}
	i, err := row.LastInsertId()
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			fmt.Println("回滚失败:", err)
			return 0, err
		}
		return 0, err
	}
	err = tx.Commit()
	if err != nil {
		fmt.Println("提交失败:", err)
		return 0, err
	}
	return i, nil
}

func SaveCommentPhoto(url string, i int64) error {
	stmt, err := dB.Prepare("insert into commentPhoto(commentId, url) values(?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(i, url)
	return err
}

func SaveCommentVideo(url string, i int64) error {
	stmt, err := dB.Prepare("insert into commentVideo(commentId, url) values(?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(i, url)
	return err
}

func ViewComment(c model.Comment) (map[int]model.Comment, error) {
	m := make(map[int]model.Comment)
	stmt, err := dB.Prepare("select commentId, gid, uid, comment, grade, IsAnonymous, time from goodsComment where gid = ? and fCommentId = -1")
	if err != nil {
		return m, err
	}
	defer stmt.Close()
	row, err := stmt.Query(c.GId)
	if err != nil {
		return m, err
	}
	defer row.Close()
	for i := 0; row.Next(); i++ {

		err = row.Scan(&c.CommentId, &c.GId, &c.UId, &c.Comment, &c.Grade, &c.IsAnonymous, &c.Time)

		if c.IsAnonymous {
			c.Name = "匿名用户"
		}
		err = dB.QueryRow("select url from commentVideo where commentId = ?", c.CommentId).Scan(&c.VideoUrl)
		if err != nil {
			if err.Error()[4:] != " no rows in result set" {
				return m, err
			}
		}
		rows, err1 := dB.Query("select url from commentPhoto where commentId = ?", c.CommentId)
		if err1 != nil {
			if err1.Error()[4:] != " no rows in result set" {
				return m, err
			}
		}
		if err1 == nil {
			for k := 0; rows.Next(); k++ {
				var url string
				err = rows.Scan(&url)
				if err != nil {
					return m, err1
				}
				c.PhotoUrl[k] = url
			}
			rows.Close()
		}

		err = dB.QueryRow("select name from User where uid = ?", c.UId).Scan(&c.Name)
		if err != nil {
			return m, err
		}
		m[i] = c

		if err != nil {
			return m, err
		}
	}
	return m, nil
}

func ReplyComment(c model.Comment) error {
	tx, err := dB.Begin()
	if err != nil {
		err1 := tx.Rollback()
		if err1 != nil {
			return err
		}
		return err
	}
	stmt, err := tx.Prepare("insert into goodsComment(gid,uid, comment, IsAnonymous, time, fCommentId) values(?,?,?,?,?,?)")
	if err != nil {
		err1 := tx.Rollback()
		if err1 != nil {
			return err
		}
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(c.GId, c.UId, c.Comment, c.IsAnonymous, time.Now(), c.CommentId)
	if err != nil {
		err1 := tx.Rollback()
		if err1 != nil {
			return err
		}
		return err
	}
	_, err = dB.Exec("update goodsComment set commentAmount = commentAmount +1 where commentId = ?", c.CommentId)
	if err != nil {
		err1 := tx.Rollback()
		if err1 != nil {
			return err
		}
		return err
	}
	err = tx.Commit()
	return err
}

func ViewSpecificComment(c model.Comment) (model.Comment, error) {
	stmt, err := dB.Prepare("select commentId from goodsComment where fCommentId = ?")
	if err != nil {
		return c, err
	}
	defer stmt.Close()

	row, err := stmt.Query(c.CommentId)
	if err != nil {
		return c, err
	}
	defer row.Close()
	for i := 0; row.Next(); i++ {
		err = row.Scan(&c.CommentId)
		if err != nil {
			return c, err
		}

		//一个赋值了commentId的model.comment,和一个空的装评论的map
		c, err = ViewSonComment(c, 0)
		if err != nil {
			if err.Error()[4:] == " no rows in result set" {
				continue
			}
			return c, err
		}
	}
	return c, nil
}

func ViewSonComment(c model.Comment, i int) (model.Comment, error) {
	//创建一个新的SComment
	sc := model.Comment{}
	//查询所有的评论信息
	stmt, err := dB.Prepare("select commentId, uid, comment, grade, IsAnonymous, time from goodsComment where fCommentId = ?")
	if err != nil {
		return sc, err
	}
	defer stmt.Close()
	row, err := stmt.Query(c.CommentId)
	if err != nil {
		return sc, err
	}
	defer row.Close()
	for row.Next() {
		//赋值所有信息
		err = row.Scan(&sc.CommentId, &sc.UId, &sc.Comment, &sc.Grade, &sc.IsAnonymous, &sc.Time)
		if err != nil {
			return sc, err
		}
		c.SComment[i] = sc

		i++

		sc, err = ViewSonComment(c, i)
		if err != nil {
			if err.Error()[4:] == " no rows in result set" {
				continue
			}
			return c, err
		}
	}
	return c, nil
}
