package main

import (
	"fmt"
	"math/rand"
	"net"
	"net/rpc"
	"sync"
	"time"
)

// m成员是一个map类型，用于存储KV数据。
// filter成员对应每个Watch调用时定义的过滤器函数列表。通知函数
// 而mu成员为互斥锁，用于在多个Goroutine访问或修改时对其它成员提供保护。
type KVStoreService struct {
	m      map[string]string
	filter map[string]func(key string)
	mu     sync.Mutex
}

func (p *KVStoreService) Get(key string, value *string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if v, ok := p.m[key]; ok {
		*value = v
		return nil
	}

	return fmt.Errorf("not found")
}

// Set方法中，输入参数是key和value组成的数组，用一个匿名的空结构体表示忽略了输出参数。当修改某个key对应的值时会调用每一个过滤器函数。
func (p *KVStoreService) Set(kv [2]string, reply *struct{}) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	key, value := kv[0], kv[1]

	if oldValue := p.m[key]; oldValue != value {
		for _, fn := range p.filter { // 执行所有的通知函数
			fn(key)
		}
	}

	p.m[key] = value
	return nil
}

// Watch方法的输入参数是超时的秒数。当有key变化时将key作为返回值返回。
// 如果超过时间后依然没有key被修改，则返回超时的错误。
// Watch的实现中，用唯一的id表示每个Watch调用，然后根据id将自身对应的过滤器函数注册到p.filter列表。
func (p *KVStoreService) Watch(timeoutSecond int, keyChanged *string) error {
	id := fmt.Sprintf("watch-%s-%03d", time.Now(), rand.Int())
	ch := make(chan string, 10) // buffered

	p.mu.Lock()
	p.filter[id] = func(key string) { ch <- key } // 设置通知函数
	p.mu.Unlock()

	select {
	case <-time.After(time.Duration(timeoutSecond) * time.Second):
		return fmt.Errorf("timeout")
	case key := <-ch:
		*keyChanged = key
		return nil
	}

	return nil
}

func NewKVStoreService() *KVStoreService {
	return &KVStoreService{
		m:      make(map[string]string),
		filter: make(map[string]func(key string)),
	}
}

func main() {
	rpc.RegisterName("KVStoreService", NewKVStoreService())
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("ListenTCP error:", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept error:", err)
		}
		rpc.ServeConn(conn)
	}
}
