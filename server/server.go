package main

import "fmt"
import kcp "github.com/xtaci/kcp-go"

func main() {
	fmt.Println("vim-go")
	lis, err := kcp.ListenWithOptions("127.0.0.1:2000", nil, 10, 3)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	cli, err := lis.AcceptKCP()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var buf = make([]byte, 1024)
	n, err := cli.Read(buf)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(buf[0:n]))
}
