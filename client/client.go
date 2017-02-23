package main

import (
	"fmt"
	"time"

	kcp "github.com/xtaci/kcp-go"
)

func main() {
	fmt.Println("vim-go")
	client, err := kcp.DialWithOptions("127.0.0.1:2000", nil, 10, 3)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	n, err := client.Write([]byte("hello"))
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
	fmt.Printf("%d\n", n)
	<-time.After(2 * time.Second)
}
