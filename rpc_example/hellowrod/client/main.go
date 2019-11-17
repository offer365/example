package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7890")
	if err != nil {
		fmt.Println("dialing:", err)
	}

	var reply = "world"

	for i := 0; i < 10; i++ {
		err = client.Call("HelloService.Hello", "hello", &reply)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(reply)
	}

}
