package main

import (
	"net"
	"reflect"
	"rpc/rpc"
)

type Client struct {
	conn net.Conn
}

func NewClient(conn net.Conn) *Client{
	return &Client{conn:conn}
}

func (c *Client) callRPC(rpcName string, fPtr interface{}) {
	fn := reflect.ValueOf(fPtr).Elem()

	f := func(args []reflect.Value)[]reflect.Value{
		inArgs  := make([]interface{}, 0, len(args))
		for _, arg := range  args {
			inArgs = append(inArgs, arg.Interface())
		}
		//连接
		cliSession := rpc.NewSession(c.conn)
		reqRPC := rpc.RPCData{Name:rpcName, Args:inArgs}
		b, err := rpc.Encode(reqRPC)
		if err != nil {
			panic(err)
		}
		//写数据
		err = cliSession.Wirte(b)
		if err != nil {
			panic(err)
		}

		respBytes, err := cliSession.Read()
		if err != nil {
			panic(err)
		}
		respRPC, err := rpc.Decode(respBytes)
		if err != nil {
			panic(err)
		}
		outArgs := make([]reflect.Value, 0, len(respRPC.Args))
		for i, arg := range  respRPC.Args {

			if arg == nil {
				outArgs = append(outArgs, reflect.Zero(fn.Type().Out(i)))
				continue
			}

		}
		return  outArgs

	}

	v := reflect.MakeFunc(fn.Type(), f)
	fn.Set(v)

}
