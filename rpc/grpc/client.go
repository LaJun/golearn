package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc/message"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:8090", grpc.WithInsecure())
	fmt.Println(conn, err)
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}
	defer conn.Close()

	orderServiceClient := message.NewOrderServiceClient(conn)
	orderRequest := &message.OrderRequest{OrderId: "201907310002", TimeStamp: time.Now().Unix()}
	orderInfo, err := orderServiceClient.GetOrderInfo(context.Background(), orderRequest)
	if orderInfo != nil {
		fmt.Println(orderInfo, orderInfo.OrderName, orderInfo.GetOrderName())
	}
}