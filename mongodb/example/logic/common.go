package logic

import (
	"context"
	"encoding/json"
	dao "github.com/offer365/example/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"io/ioutil"
	"time"
)

var db dao.DB
var store dao.Store

func Init(host, port, user, pwd, database string, timeout time.Duration, ci map[string]string) (err error) {
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	db = dao.NewDB("mongo")
	err = db.Init(ctx,
		dao.WithHost(host),
		dao.WithPort(port),
		dao.WithUsername(user),
		dao.WithPwd(pwd),
		dao.WithDatabase(database),
		dao.WithTimeout(timeout),
		dao.WithCollIndex(ci),
	)
	if err != nil {
		return
	}
	store = dao.NewStore("mongo")
	return store.Init(
		ctx,
		dao.WithHost(host),
		dao.WithPort(port),
		dao.WithUsername(user),
		dao.WithPwd(pwd),
		dao.WithDatabase(database),
		dao.WithTimeout(timeout),
	)
}

func Update(coll string, id string, body io.Reader) (err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", oid}}
	byt, err := ioutil.ReadAll(body)
	if err != nil {
		return
	}
	data := make(map[string]interface{}, 0)
	err = json.Unmarshal(byt, &data)
	if err != nil {
		return
	}
	update := bson.D{}
	for k, v := range data {
		update = append(update, bson.E{"$set", bson.D{{k, v}}})
	}
	return db.Update(coll, filter, update)
}

func Delete(coll string, id string) (err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", oid}}
	return db.Delete(coll, filter)
}

func Count(coll string) (num int64, err error) {
	return db.Count(coll, make(map[string]string))
}
