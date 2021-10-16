package main

import (
	"fmt"
	"net"
	"reflect"
	"rpc/rpc"
)

type Server struct {
	addr string
	funcs map[string]reflect.Value
}

func NewServer(addr string) * Server{
	return &Server{addr: addr, funcs: make(map[string]reflect.Value)}
}

func (s *Server) Register (rpcName string, f interface{}){
	if _,ok := s.funcs[rpcName]; ok {
		return
	}
	fVal := reflect.ValueOf(f)
	s.funcs[rpcName] = fVal
}

func (s * Server) Run() {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		fmt.Printf("监听 %s err: %v", s.addr, err)
		return
	}
	for {
		conn, err := lis.Accept()
		if err != nil {
			return
		}
		serSession := rpc.NewSession(conn)
		b, err := serSession.Read()
		if err != nil {
			return
		}
		rpcData, err := rpc.Decode(b)
		if err != nil {
			return
		}
		f, ok := s.funcs[rpcData.Name]
		if !ok {
			fmt.Println("函数 %s 不存在", rpcData.Name)
			return
		}
		inArgs := make([]reflect.Value,0, len(rpcData.Args))
		for _, arg := range rpcData.Args {
			inArgs = append(inArgs, reflect.ValueOf(arg))
		}
		out := f.Call(inArgs)
		outArgs := make([]interface{}, 0, len(out))
		for _, o := range out {
			outArgs = append(outArgs, o.Interface())
		}

		respRPCData := rpc.RPCData{rpcData.Name, outArgs}
		bytes, err := rpc.Encode(respRPCData)
		if err != nil{
			return
		}
		err = serSession.Wirte(bytes)
		if err != nil {
			return
		}
	}




}
