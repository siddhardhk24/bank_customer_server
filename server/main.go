package main

import (
	"context"
	"fmt"
	"net"

	"github.com/siddhardhk24/bank_customer_server/controllers"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	// pro "github.com/grpc_bank/bank_customer_proto/netxd_customer"
)

func initDatabase(client *mongo.Client) {
	profileCollection := config.GetCollection(client, "bankdb", "billa")
	controllers.CustomerService = services.InitCustomerService(profileCollection, context.Background())
}

func main() {
	mongoclient, err := config.ConnectDataBase()
	defer mongoclient.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	initDatabase(mongoclient)
	lis, err := net.Listen("tcp", constants.Port)
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	pro.RegisterCustomer_ServiceServer(s, &controllers.RPCServer{})

	fmt.Println("Server listening on", constants.Port)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}
}
