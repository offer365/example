package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/url"
)

func main()  {
	u:=url.URL{Scheme:"ws",Host:"10.0.0.92",Path:"realtime"}
	c,_,err:=websocket.DefaultDialer.Dial(u.String(),nil)
	fmt.Println("err:",err)
	defer c.Close()
	err=c.WriteMessage(websocket.BinaryMessage,[]byte("1"))
	a,b,err:=c.ReadMessage()
	fmt.Println(a,string(b),err)
}
