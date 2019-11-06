package mongodb

import (
	"context"
	"fmt"
	"testing"
	"time"
)

var db DB
var store Store

func TestNewDB(t *testing.T) {
	db = NewDB("mongo")
	err := db.Init(context.Background(),
		WithDB("example"),
		WithPort("27017"),
		WithUsername("example"),
		WithPwd("example"),
		WithHost("10.0.0.200"),
		WithTimeout(time.Second*2))
	fmt.Println(err)
}

func TestNewStore(t *testing.T) {
	store = NewStore("mongo")
	err := store.Init(context.Background(),
		WithDB("example"),
		WithPort("27017"),
		WithUsername("example"),
		WithPwd("example"),
		WithHost("10.0.0.200"),
		WithTimeout(time.Second*2))
	fmt.Println(err)
}
