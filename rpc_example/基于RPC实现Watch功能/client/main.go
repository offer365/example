package main

import (
	"fmt"
	"log"
	"net/rpc"
	"strconv"
	"time"
)

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println("dialing:", err)
	}

	doClientWork(client)

}

func doClientWork(client *rpc.Client) {
	go func() {
		var keyChanged string
		for range time.NewTicker(time.Second).C {
			err := client.Call("KVStoreService.Watch", 30, &keyChanged)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("watch:", keyChanged)
		}

	}()

	for range time.NewTicker(time.Second).C {
		err := client.Call(
			"KVStoreService.Set", [2]string{"abc", strconv.Itoa(int(time.Now().UnixNano()))},
			new(struct{}),
		)
		if err != nil {
			log.Fatal(err)
		}
	}

	// time.Sleep(time.Second*13)
}
