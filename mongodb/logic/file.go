package logic

import (
	"context"
	"github.com/offer365/example/mongodb/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
	"time"
)

func Upload(name string, source io.Reader) (id string, err error) {
	return store.Upload(name, source)
}

func Download(id string, source io.Writer) (size int64, err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	return store.Download(oid, source)
}

func FindFile(id string, skip, limit int32) (instances []*model.File, err error) {
	var (
		oid    primitive.ObjectID
		filter interface{}
	)
	if id == "" {
		filter = make(map[string]string)
	} else {
		if oid, err = primitive.ObjectIDFromHex(id); err != nil {
			return
		}
		filter = bson.D{{"_id", oid}}
	}

	fu := func(cursor *mongo.Cursor) (err error) {
		// 遍历结果集
		ctx, _ := context.WithTimeout(context.Background(), time.Second*2)
		for cursor.Next(ctx) {
			instance := new(model.File)
			if err = cursor.Decode(instance); err == nil { // 反序列化bson到对象
				instances = append(instances, instance)
			}
		}
		return
	}
	err = store.FindFile(filter, fu, skip, limit)
	return
}

func DeleteFile(id string) (err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	return store.DeleteFile(oid)
}
