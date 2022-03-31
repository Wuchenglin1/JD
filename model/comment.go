package model

import (
	"gorm.io/gorm"
	"mime/multipart"
	"time"
)

type Comment struct {
	gorm.Model
	UId         uint
	Name        string
	GId         int
	Comment     string
	Grade       int
	VideoUrl    string         `gorm:"Column:videoUrl"`
	PhotoUrl    map[int]string `gorm:"Column:photoUrl"`
	IsAnonymous bool
	Time        time.Time
	FCommentId  uint
}

type CommentFile struct {
	Video multipart.File
	Photo map[string]multipart.File
}
