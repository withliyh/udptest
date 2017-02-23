package main

import (
	"flag"
	"fmt"
	"time"

	kcp "github.com/xtaci/kcp-go"
)

var host = flag.String("host", "127.0.0.1:2000", "remote host address and port")

func main() {
	flag.Parse()
	client, err := kcp.DialWithOptions(*host, nil, 10, 3)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("connect to %s.\n", *host)

	content := "hello"
	n, err := client.Write([]byte(content))
	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	}
	fmt.Printf("send [%s] to target [%d] bytes\n", content, n)
	var buf = make([]byte, 1024)
	n, err = client.Read(buf)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	}
	fmt.Printf("rect [%s]\n", string(buf[0:n]))
	<-time.After(10 * time.Second)
}
