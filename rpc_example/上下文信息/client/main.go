package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println("dialing:", err)
	}

	var reply = ""

	err = client.Call("HelloService.Login", "user:password", &reply)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reply)

	client.Call("HelloService.Hello", "mali", &reply)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reply)

}
