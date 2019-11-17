package logic

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"time"

	"../model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindOnePrincipal(coll string, id string) (instances []*model.Principal, err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", oid}}
	instance := new(model.Principal)
	err = db.FindOne(coll, filter, instance)
	return []*model.Principal{instance}, err
}

func FindAllPrincipal(coll string, skip, limit int64) (instances []*model.Principal, err error) {
	instances = make([]*model.Principal, 0)
	fu := func(cursor *mongo.Cursor) (err error) {
		// 遍历结果集
		ctx, _ := context.WithTimeout(context.Background(), time.Second*2)
		for cursor.Next(ctx) {
			instance := new(model.Principal)
			if err = cursor.Decode(instance); err == nil { // 反序列化bson到对象
				instances = append(instances, instance)
			}
		}
		return
	}
	err = db.Find(coll, make(map[string]string), fu, skip, limit, -1)
	return
}

func InsertPrincipal(coll string, body io.ReadCloser) (id string, err error) {
	byt, err := ioutil.ReadAll(body)
	if err != nil {
		return
	}
	instance := new(model.Principal)
	err = json.Unmarshal(byt, instance)
	if err != nil {
		return
	}
	instance.ID = primitive.NewObjectID()
	return db.Insert(coll, instance)
}
