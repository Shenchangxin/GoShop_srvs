package main

import (
	"google.golang.org/grpc"
	"mxshop_srvs/goods_srv/proto"
)

var brandClient proto.GoodsClient
var conn *grpc.ClientConn

func Init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	brandClient = proto.NewGoodsClient(conn)
}

func main() {
	Init()

	TestGetCategoryList()
	conn.Close()
}
