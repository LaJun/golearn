package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc2/message"
	"io"
	"net"
)
type OrderServiceImpl struct {

}

func (os *OrderServiceImpl) GetOrderInfos(stream message.OrderService_GetOrderInfosServer) error{
	fmt.Println("服务端流模式")
	for {
		orderRequest, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("数据读取结束")
			return err
		}
		if err != nil {
			panic(err.Error())
			return err
		}
		fmt.Println(orderRequest.GetOrderId())
		orderMap := map[string]message.OrderInfo{
			"201907300001": message.OrderInfo{OrderId: "201907300001", OrderName: "衣服", OrderStatus: "已付款"},
			"201907310001": message.OrderInfo{OrderId: "201907310001", OrderName: "零食", OrderStatus: "已付款"},
			"201907310002": message.OrderInfo{OrderId: "201907310002", OrderName: "食品", OrderStatus: "未付款"},
		}

		result := orderMap[orderRequest.GetOrderId()]
		err = stream.Send(&result)
		if err == io.EOF {
			fmt.Println(err)
			return err
		}
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		return nil
	}
}

func main() {
	server := grpc.NewServer()
	//注册
	message.RegisterOrderServiceServer(server, new(OrderServiceImpl))
	lis, err := net.Listen("tcp", ":8881")
	if err != nil {
		panic(err.Error())
	}
	server.Serve(lis)
}