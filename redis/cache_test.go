package redis

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

var conn redis.Conn
var rdb Cache

func TestNewCache(t *testing.T) {
	rdb = NewCache("redis")
	rdb.Init(
		WithAddr("10.0.0.92:6379"))
	conn = rdb.Conn()
	defer conn.Close()
	// 文档  https://godoc.org/github.com/gomodule/redigo/redis#pkg-examples
	// 命令  http://redisdoc.com
	//ExampleKV()
	//ExampleList()
	//ExampleHash()
	ExampleJson()
	ExampleMulti()
	//ExampleExpire()
	//ExamplePip()
	//ExamplePublish()
}

// 普通的键值操作
func ExampleKV() {
	_, err := conn.Do("set", "a", "11")
	if err != nil {
		fmt.Println("set error,", err)
	}
	str, err := redis.String(conn.Do("get", "a"))
	fmt.Println(str, err)

	_, err = conn.Do("Set", "a", 100)
	if err != nil {
		fmt.Println("set error,", err)
	}
	val, err := redis.Int(conn.Do("Get", "a"))
	fmt.Println(val, err)

	// SET的附加参数 指定过期时间
	_, err = conn.Do("Set", "aaa", 100, "ex", 5)
	if err != nil {
		fmt.Println("set error,", err)
	}
	for i := 0; i < 6; i++ {
		time.Sleep(time.Second)
		val, err = redis.Int(conn.Do("Get", "aaa"))
		fmt.Println(val, err)
	}

	//批量写入读取
	//MGET key [key …]
	//MSET key value [key value …]
	// mset 同时设置多个键值对
	_, err = conn.Do("MSet", "aa", 102, "bb", 103, "cc", 104)
	if err != nil {
		fmt.Println("hset error,", err)
	}
	rs, err := redis.Ints(conn.Do("MGet", "aa", "bb", "cc"))
	fmt.Println(rs)
	for _, v := range rs {
		fmt.Println(v)
	}

	values, err := redis.Values(conn.Do("MGet", "aa", "bb", "cc"))
	var v1 int
	for len(values) > 0 {
		values, err = redis.Scan(values, &v1)
		fmt.Println(v1)
	}

	var value1 int
	var value2 string
	reply, err := redis.Values(conn.Do("MGET", "key1", "key2"))
	if err != nil {
		// handle error
	}
	if _, err := redis.Scan(reply, &value1, &value2); err != nil {
		// handle error
	}

}

func ExampleList() {
	// 队列 操作 类似于管道 1 先进，其次是2 从左边推进去
	conn.Do("del", "list1")
	_, err := conn.Do("lpush", "list1", "1", "2", "3", "4", "5")
	if err != nil {
		fmt.Println("lpush error,", err)
	}

	list, _ := redis.Strings(conn.Do("lrange", "list1", 0, 100))
	for _, v := range list {
		fmt.Println(v)
	}
	values, _ := redis.Values(conn.Do("lrange", "list1", 0, 100))
	// 循环迭代方法一
	for _, v := range values {
		fmt.Println(string(v.([]byte)))
	}

	// 循环迭代方法二
	var v1 string
	for len(values) > 0 {
		values, err = redis.Scan(values, &v1)
		fmt.Println(v1)
	}

	// 从左边拉出来
	str, err := redis.String(conn.Do("lpop", "list1"))
	fmt.Println(str)
	// 从右边拉出来
	str, err = redis.String(conn.Do("rpop", "list1"))
	fmt.Println(str)
}

func ExampleHash() {
	// 设置哈希表
	_, err := conn.Do("HSet", "hash1", "a", 101)
	if err != nil {
		fmt.Println("hset error,", err)
	}
	r, err := redis.Int(conn.Do("HGet", "hash1", "a"))
	fmt.Println(r)

	//批量写入读取对象(Hashtable)
	//HMSET key field value [field value …]
	//HMGET key field [field …]
	_, err = conn.Do("HMSet", "hash1", "a", 101, "b", 101, "c", 102, "d", 103)
	if err != nil {
		fmt.Println("hset error,", err)
	}

	rs, err := redis.Ints(conn.Do("HMGet", "hash1", "a", "b", "c", "d"))
	fmt.Println(rs)
}

func ExampleJson() {
	imap := map[string]string{"name": "mali", "age": "22"}
	value, _ := json.Marshal(imap)
	// SETNX 是『SET if Not eXists』(如果不存在，则 SET)的简写。
	n, err := conn.Do("SETNX", "user1", value)
	if err != nil {
		fmt.Println(err)
	}
	if n == int64(1) {
		fmt.Println("success")
	}

	var imapGet map[string]string

	valueGet, err := redis.Bytes(conn.Do("GET", "user1"))
	if err != nil {
		fmt.Println(err)
	}

	errShal := json.Unmarshal(valueGet, &imapGet)
	if errShal != nil {
		fmt.Println(err)
	}
	fmt.Println(imapGet["name"])
	fmt.Println(imapGet["age"])
}

func ExampleExpire() {
	// 设置超时时间  expire 关键字 a 这个键，10 超时 10 秒
	_, err := conn.Do("expire", "a", 3)
	if err != nil {
		fmt.Println("hset error,", err)
	}
	var exist bool = true
	for exist {
		exist, err = redis.Bool(conn.Do("exists", "a"))
		fmt.Println("exist:", exist)
		time.Sleep(time.Second)
	}
}

// 管道
func ExamplePip() {
	conn.Send("set", "name", "lisi") //将命令写入缓冲区。
	conn.Send("get", "name")
	conn.Flush()                              //将缓冲区的内容刷新到服务器。
	fmt.Println(redis.String(conn.Receive())) // 从服务器读取单个答复
	fmt.Println(redis.String(conn.Receive()))

	// 流水线事务
	conn.Send("MULTI")
	conn.Send("INCR", "foo")
	conn.Send("INCR", "bar")
	r, err := conn.Do("EXEC")
	fmt.Println(r, err) // prints [1, 1]
}

//事务
func ExampleMulti() {

	//MULTI：开启事务
	//
	//EXEC：执行事务
	//
	//DISCARD：取消事务
	//
	//WATCH：监视事务中的键变化，一旦有改变则取消事务。

	// 流水线事务
	conn.Send("MULTI")
	conn.Send("INCR", "foo")
	conn.Send("INCR", "bar")
	r, err := conn.Do("EXEC")
	fmt.Println(r, err) // prints [1, 1]
}

// 订阅
func ExampleSubscribe() {
	c := rdb.Conn()
	defer c.Close()
	psc := redis.PubSubConn{c}
	psc.Subscribe("redChatRoom")
	//psc.Unsubscribe("redChatRoom")  // 退订
	for {
		switch v := psc.Receive().(type) {
		case redis.Message: // 接受到的信息
			fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
		case redis.Subscription: // 这是执行订阅的反馈信息
			fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			fmt.Println(v)
			return
		}
	}
}

// 发布
func ExamplePublish() {
	go ExampleSubscribe()
	go ExampleSubscribe()
	go ExampleSubscribe()
	go ExampleSubscribe()
	go ExampleSubscribe()
	go ExampleSubscribe()
	c := rdb.Conn()
	defer c.Close()
	for {
		var s string
		rand.Seed(time.Now().UnixNano())
		s = strconv.Itoa(rand.Int())
		fmt.Scanln(&s)
		_, err := c.Do("PUBLISH", "redChatRoom", s)
		if err != nil {
			fmt.Println("pub err: ", err)
			return
		}
		time.Sleep(time.Second)
	}
}
