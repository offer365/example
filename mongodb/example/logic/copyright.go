package logic

import (
	"../model"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
	"io/ioutil"
	"time"
)

func FindOneCopyright(coll string, id string) (instances []*model.Copyright, err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", oid}}
	instance := new(model.Copyright)
	err = db.FindOne(coll, filter, instance)
	return []*model.Copyright{instance}, err
}

func FindAllCopyright(coll string, skip, limit int64) (instances []*model.Copyright, err error) {
	instances = make([]*model.Copyright, 0)
	fu := func(cursor *mongo.Cursor) (err error) {
		// 遍历结果集
		ctx, _ := context.WithTimeout(context.Background(), time.Second*2)
		for cursor.Next(ctx) {
			instance := new(model.Copyright)
			if err = cursor.Decode(instance); err == nil { // 反序列化bson到对象
				instances = append(instances, instance)
			}
		}
		return
	}
	err = db.Find(coll, make(map[string]string), fu, skip, limit, -1)
	return
}

func InsertCopyright(coll string, body io.Reader) (id string, err error) {
	byt, err := ioutil.ReadAll(body)
	if err != nil {
		return
	}
	instance := new(model.Copyright)
	err = json.Unmarshal(byt, instance)
	if err != nil {
		return
	}
	instance.ID = primitive.NewObjectID()
	return db.Insert(coll, instance)
}
