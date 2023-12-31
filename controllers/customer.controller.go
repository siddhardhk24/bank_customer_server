package controllers

import (
	"context"

	pro "github.com/siddhardhk24/bank_customer_proto/netxd_customer"
	"github.com/siddhardhk24/bank_customer_service/interfaces"
	"github.com/siddhardhk24/bank_customer_service/models"
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
