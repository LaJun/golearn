package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"testing"
)

type User struct {
	Name string
	Age int
}

func queryUser(uid int) (User, error) {
	user := make(map[int]User)
	//假数据
	user[0] = User{"zs", 20}
	user[1] = User{"lisi", 21}
	//模拟查询用户
	if u, ok := user[uid]; ok{
		return u, nil
	}
	return User{}, fmt.Errorf("%d err", uid)
}

func TestRPC(t *testing.T)  {
	gob.Register(User{})
	addr := "127.0.0.1:8000"
	srv := NewServer(addr)
	srv.Register("queryUser", queryUser)
	go srv.Run()
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println(err)
	}
	cli := NewClient(conn)
	var query func(int) (User,error)
	cli.callRPC("queryUser", &query)
	u, err := queryUser(1)
	if err != nil {
		fmt.Println("err")
	}
	fmt.Println(u)
}

