package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Params struct {
	Width, Height int
}

type Rect struct {}

func (r *Rect)Area(p Params, operator string, ret *int)(error) {
	switch operator {
	case "*":* ret = p.Height * p.Width
	default:
		*ret = p.Height + p.Width
	}

	return nil
}
func main()  {
	rect := new(Rect)
	rpc.Register(rect)
	rpc.HandleHTTP()
	lis, err := net.Listen("tcp", ":8800")
	if err != nil {
		log.Panicln(err)
	}
	for {
		conn, err := lis.Accept()
		if err != nil {
			continue
		}
		go func(conn net.Conn) {
			fmt.Println("new client")
			jsonrpc.ServeConn(conn)
		}(conn)
	}
}