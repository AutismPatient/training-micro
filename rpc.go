package main

import (
	"fmt"
	"log"
	"net/http"
	"net/rpc"
)

type User struct {
	Name string
	Age  int
}

func (u *User) Print(ages *User, reply *string) error {
	*reply = fmt.Sprintf("name: %s,age: %d", ages.Name, ages.Age)
	return nil
}

// http 形式
func runRpc() error {
	err := rpc.RegisterName("test_rpc", new(User))
	if err != nil {
		return err
	}
	rpc.HandleHTTP()

	if err = http.ListenAndServe(":8085", nil); err != nil {
		log.Fatal(err.Error())
	}
	return nil
}

func main() {
	err := runRpc()
	if err != nil {
		panic(err)
	}
}
