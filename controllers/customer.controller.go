package controllers

import (
	"context"

	pro "github.com/grpc_bank/bank_customer_proto/netxd_customer"
)

type RPCServer struct {
	pro.UnimplementedCustomer_ServiceServer
}

var (
	CustomerService interfaces.ICustomer
)

func (s *RPCServer) CreateCustomer(ctx context.Context, req *pro.Customer) (*pro.CustomerResponse, error) {
	dbCustomer := &models.Customer{CustomerId: req.CustomerId}
	result, err := CustomerService.CreateCustomer(dbCustomer)
	if err != nil {
		return nil, err
	} else {
		responseCustomer := &pro.CustomerResponse{
			CustomerId: result.CustomerId,
		}
		return responseCustomer, nil
	}
}
