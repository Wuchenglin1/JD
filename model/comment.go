package model

import (
	"mime/multipart"
	"time"
)

type Comment struct {
	UId       int    `json:"uid"`
	Name      string `json:"name"`
	GId       int    `json:"gid"`
	CommentId int    `json:"commentId"`
	Comment   string `json:"comment"`
	//CommentInt int                    `json:"sonCommentInt"` //子评论数
	//Praise        int                    `json:"praise"`        //点赞数
	Grade       int                       `json:"grade"`
	Video       multipart.File            `json:"video"`
	VideoUrl    string                    `json:"videoUrl"`
	Photo       map[string]multipart.File `json:"photo"`
	PhotoUrl    map[int]string            `json:"photoUrl"`
	IsAnonymous bool                      `json:"isAnonymous"`
	Time        time.Time                 `json:"time"`
	SComment    map[int]Comment           `json:"sComment"`
}
