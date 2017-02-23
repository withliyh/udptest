package main

import "fmt"
import kcp "github.com/xtaci/kcp-go"

func main() {
	lis, err := kcp.ListenWithOptions("0.0.0.0:2000", nil, 10, 3)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("0.0.0.0:2000 is listening.")
	for {
		cli, err := lis.AcceptKCP()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			go handle(cli)
		}
	}
}

func handle(client *kcp.UDPSession) {
	var buf = make([]byte, 1024)
	defer func() {
		client.Close()
		fmt.Printf("close connect [%s]\n", client.RemoteAddr().String())
	}()
	for {
		n, err := client.Read(buf)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		remoteAddr := client.RemoteAddr().String()
		fmt.Printf("rect [%s] from [%s]\n", string(buf[0:n]), remoteAddr)

		n, err = client.Write(buf[0:n])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("send [%s] to [%s]\n", string(buf[0:n]), remoteAddr)
	}
}
