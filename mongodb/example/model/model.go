package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// 产品
type Product struct {
	ID           primitive.ObjectID   `bson:"_id" json:"id"`
	Name         string               `bson:"name" json:"name" index:"unique"`
	Number       string               `bson:"number" json:"number"`                               // 产品编号
	Date         time.Time            `bson:"date" json:"date"`                                   // 发布日期
	Version      string               `bson:"version" json:"version"`                             // 版本号
	Introduction string               `bson:"introduction" json:"introduction"`                   // 简介
	Principal    []primitive.ObjectID `bson:"principal" json:"principal" collection:"principals"` //负责人
	Projects     []primitive.ObjectID `bson:"projects" json:"projects" collection:"projects"`     // 项目
	Copyright    []primitive.ObjectID `bson:"copyright" json:"copyright" collection:"copyrights"` //软著
	Files        []primitive.ObjectID `bson:"files" json:"files"`
}

// 用户表
type Principal struct {
	ID    primitive.ObjectID `bson:"_id" json:"id"`
	Name  string             `bson:"name" index:"unique" json:"name"`
	Email string             `bson:"email" json:"email"`
}

// 项目
type Project struct {
	ID      primitive.ObjectID `bson:"_id" json:"id"`
	Name    string             `bson:"name" index:"unique" json:"name"`     // 项目名
	Owner   string             `bson:"owner" json:"owner"`                  // 对方单位
	Contact string             `bson:"contact" json:"contact"`              // 联系人
	Number  string             `bson:"number" index:"unique" json:"number"` // 项目编号
}

//软著
type Copyright struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Certificate string             `bson:"certificate" index:"unique" json:"certificate"` // 证书号
	Name        string             `bson:"name" index:"unique" json:"name"`               // 名称
	Number      string             `bson:"number" index:"unique" json:"number"`           // 编号
	Register    string             `bson:"register" index:"unique" json:"register"`       // 注册号
}

type File struct {
	ID         primitive.ObjectID `bson:"_id" json:"id"`
	FileName   string             `bson:"filename" json:"name"`
	Length     int64              `bson:"length" json:"length"`
	ChunkSize  int64              `bson:"chunkSize" json:"chunk_size"`
	UploadDate primitive.DateTime `bson:"uploadDate" json:"upload_date"`
}
