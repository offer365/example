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

type product struct {
	ID           primitive.ObjectID `bson:"_id" json:"id"`
	Name         string             `bson:"name" json:"name"`
	Number       string             `bson:"number" json:"number"`
	Version      string             `bson:"version" json:"version"`
	Date         time.Time          `bson:"date" json:"date"`
	Principal    []*model.Principal `bson:"principal" json:"principal"` // 负责人
	Projects     []*model.Project   `bson:"projects" json:"projects"`
	Copyright    []*model.Copyright `bson:"copyright" json:"copyright"`
	Files        []*model.File      `bson:"files" json:"files"`
	Introduction string             `bson:"introduction" json:"introduction"` // 简介
}

type show struct {
	ID           primitive.ObjectID `bson:"_id" json:"id"`
	Name         string             `bson:"name" json:"name"`
	Number       string             `bson:"number" json:"number"`
	Version      string             `bson:"version" json:"version"`
	Date         time.Time          `bson:"date" json:"date"`
	Principal    string             `bson:"principal" json:"principal"` // 负责人
	Projects     []string           `bson:"projects" json:"projects"`
	Copyright    string             `bson:"copyright" json:"copyright"`
	Files        []string           `bson:"files" json:"files"`
	Introduction string             `bson:"introduction" json:"introduction"` // 简介
}

func Aggregation(coll string, id string, skip, limit int64) (data interface{}, err error) {
	var pipe mongo.Pipeline
	//{"$limit",limit},
	//{"$skip",skip},
	//{"$sort",-1},
	if id != "" {
		if oid, err := primitive.ObjectIDFromHex(id); err == nil {
			o1 := bson.D{
				{"$match", bson.M{"_id": oid}},
			}
			pipe = append(pipe, o1)
		}
	}

	// db.products.aggregate([{$match:{"_id":ObjectId("5d5d0a3a306ba203ca7447a1")}},{$lookup:{from:"projects",localField:"projects",foreignField:"_id",as:"projects"}},{$lookup:{from:"users",localField:"authors",foreignField:"_id",as:"authors"}}]).pretty()
	o2 := bson.D{{
		"$lookup", bson.M{
			"from":         "projects",
			"localField":   "projects",
			"foreignField": "_id", "as": "projects",
		},
	}}
	pipe = append(pipe, o2)
	o3 := bson.D{{
		"$lookup", bson.M{
			"from":         "principals",
			"localField":   "principal",
			"foreignField": "_id", "as": "principal",
		},
	}}
	pipe = append(pipe, o3)
	o4 := bson.D{{
		"$lookup", bson.M{
			"from":         "fs.files",
			"localField":   "files",
			"foreignField": "_id", "as": "files",
		},
	}}
	pipe = append(pipe, o4)
	o5 := bson.D{{
		"$lookup", bson.M{
			"from":         "copyrights",
			"localField":   "copyright",
			"foreignField": "_id", "as": "copyright",
		},
	}}
	pipe = append(pipe, o5)
	pipe = append(pipe, bson.D{{"$limit", limit}})
	pipe = append(pipe, bson.D{{"$skip", skip}})
	//pipe=append(pipe,bson.D{{"$sort",-1}})

	instances := make([]*product, 0)
	fu := func(cursor *mongo.Cursor) (err error) {
		// 遍历结果集
		ctx, _ := context.WithTimeout(context.Background(), time.Second*2)
		for cursor.Next(ctx) {
			instance := new(product)
			if err = cursor.Decode(instance); err == nil { // 反序列化bson到对象
				instances = append(instances, instance)
			}
		}
		return
	}

	err = db.Aggregation(coll, pipe, fu)
	return instances, err
}

// 聚合数据
func FindOneProduct(coll string, id string) (instances []*product, err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", oid}}
	instance := new(model.Product)
	err = db.FindOne(coll, filter, instance)
	principals := make([]*model.Principal, 0)
	for i, id := range instance.Principal {
		ins := new(model.Principal)
		err = db.FindOne(coll, bson.D{{"_id", id}}, ins)
		principals[i] = ins
	}
	projects := make([]*model.Project, 0)
	for i, id := range instance.Projects {
		ins := new(model.Project)
		err = db.FindOne(coll, bson.D{{"_id", id}}, ins)
		projects[i] = ins
	}
	copyrights := make([]*model.Copyright, 0)
	for i, id := range instance.Copyright {
		ins := new(model.Copyright)
		err = db.FindOne(coll, bson.D{{"_id", id}}, ins)
		copyrights[i] = ins
	}
	files := make([]*model.File, 0)
	for i, id := range instance.Files {
		ins := new(model.File)
		err = db.FindOne(coll, bson.D{{"_id", id}}, ins)
		files[i] = ins
	}
	prod := new(product)
	prod.Name = instance.Name
	prod.Number = instance.Number
	prod.Introduction = instance.Introduction
	prod.Version = instance.Version
	prod.Copyright = copyrights
	prod.Projects = projects
	prod.Principal = principals
	prod.Files = files
	return []*product{prod}, err
}

func FindAllProduct(coll string, skip, limit int64) (instances []*model.Product, err error) {
	instances = make([]*model.Product, 0)
	fu := func(cursor *mongo.Cursor) (err error) {
		// 遍历结果集
		ctx, _ := context.WithTimeout(context.Background(), time.Second*2)
		for cursor.Next(ctx) {
			instance := new(model.Product)
			if err = cursor.Decode(instance); err == nil { // 反序列化bson到对象
				instances = append(instances, instance)
			}
		}
		return
	}
	err = db.Find(coll, make(map[string]string), fu, skip, limit, -1)
	return
}

func InsertProduct(coll string, body io.Reader) (id string, err error) {
	byt, err := ioutil.ReadAll(body)
	if err != nil {
		return
	}
	pd := new(show)
	err = json.Unmarshal(byt, pd)
	if err != nil {
		return
	}
	instance := new(model.Product)

	if oid, err := primitive.ObjectIDFromHex(pd.Principal); err == nil {
		instance.Principal = append(instance.Principal, oid)
	}

	if oid, err := primitive.ObjectIDFromHex(pd.Copyright); err == nil {
		instance.Copyright = append(instance.Copyright, oid)
	}

	for _, id := range pd.Files {
		if oid, err := primitive.ObjectIDFromHex(id); err == nil {
			instance.Files = append(instance.Files, oid)
		}
	}

	for _, id := range pd.Projects {
		if oid, err := primitive.ObjectIDFromHex(id); err == nil {
			instance.Projects = append(instance.Projects, oid)
		}
	}

	instance.Name = pd.Name
	instance.Version = pd.Version
	instance.Date = time.Now().Local()
	instance.Number = pd.Number
	instance.Introduction = pd.Introduction
	if err != nil {
		return
	}
	instance.ID = primitive.NewObjectID()
	return db.Insert(coll, instance)
}

func UpdateProduct(coll string, id string, body io.Reader) (err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", oid}}
	byt, err := ioutil.ReadAll(body)
	if err != nil {
		return
	}

	pd := new(show)
	err = json.Unmarshal(byt, pd)
	if err != nil {
		return
	}
	data := make(map[string]interface{}, 0)
	principal := make([]primitive.ObjectID, 0)
	copyrights := make([]primitive.ObjectID, 0)
	projects := make([]primitive.ObjectID, 0)
	files := make([]primitive.ObjectID, 0)

	if oid, err := primitive.ObjectIDFromHex(pd.Principal); err == nil {
		principal = append(principal, oid)
	}

	if oid, err := primitive.ObjectIDFromHex(pd.Copyright); err == nil {
		copyrights = append(copyrights, oid)
	}

	for _, id := range pd.Projects {
		if oid, err := primitive.ObjectIDFromHex(id); err == nil {
			projects = append(projects, oid)
		}
	}
	for _, id := range pd.Files {
		if oid, err := primitive.ObjectIDFromHex(id); err == nil {
			files = append(files, oid)
		}
	}

	data["name"] = pd.Name
	data["version"] = pd.Version
	data["number"] = pd.Number
	data["introduction"] = pd.Introduction
	data["files"] = files
	data["projects"] = projects
	data["principal"] = principal
	data["copyright"] = copyrights
	update := bson.D{}
	for k, v := range data {
		update = append(update, bson.E{"$set", bson.D{{k, v}}})
	}
	return db.Update(coll, filter, update)
}
