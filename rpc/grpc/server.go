package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"grpc/message"
	"net"
	"time"
)

func main() {
	server := grpc.NewServer()

	message.RegisterOrderServiceServer(server, new (OrderServiceImpl))

	lis, err := net.Listen("tcp", ":8090")
	if err != nil {
		panic(err.Error())
	}
	server.Serve(lis)
}

type OrderServiceImpl struct {

}

func (os * OrderServiceImpl)GetOrderInfo(ctx context.Context, request *message.OrderRequest) (* message.OrderInfo, error)  {
	orderMap := map[string]message.OrderInfo{
		"201907300001": message.OrderInfo{OrderId: "201907300001", OrderName: "衣服", OrderStatus: "已付款"},
		"201907310001": message.OrderInfo{OrderId: "201907310001", OrderName: "零食", OrderStatus: "已付款"},
		"201907310002": message.OrderInfo{OrderId: "201907310002", OrderName: "食品", OrderStatus: "未付款"},
	}
	var response *message.OrderInfo
	current := time.Now().Unix()
	if (request.TimeStamp > current) {
		*response = message.OrderInfo{OrderId: "0", OrderName: "", OrderStatus: "订单信息异常"}
	} else {
		result := orderMap[request.OrderId]
		if result.OrderId != "" {
			fmt.Println(result)
			return &result, nil
		} else {
			return nil, errors.New("server error")
		}
	}
	return response, nil
}