package model

import (
	"mime/multipart"
	"time"
)

type Comment struct {
	UId       int       `json:"uId"`
	GId       int       `json:"gId"`
	CommentId int       `json:"commentId"`
	Comment   string    `json:"comment"`
	Color     string    `json:"color"`
	Time      time.Time `json:"time"`
	//SonCommentInt int                    `json:"sonCommentInt"` //子评论数
	//Praise        int                    `json:"praise"`        //点赞数
	Video multipart.File         `json:"video"`
	Photo map[int]multipart.File `json:"photo"`
}
