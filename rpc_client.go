package main

import (
	"fmt"
	"net/rpc"
)

type NewUser struct {
	Name string
	Age  int
}

func runClientRpc() error {
	c, err := rpc.DialHTTP("tcp", "127.0.0.1:8085")
	if err != nil {
		return err
	}
	var res string
	err = c.Call("test_rpc.Print", NewUser{
		"666",
		18,
	}, &res)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

func main() {
	err := runClientRpc()
	if err != nil {
		panic(err)
	}
}
