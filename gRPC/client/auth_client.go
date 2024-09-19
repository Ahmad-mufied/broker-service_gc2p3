package client

import (
	"context"
	"github.com/Ahmad-mufied/broker-service_gc2p3/pb"
	"google.golang.org/grpc"
	"log"
	"time"
)

type AuthClient struct {
	service     pb.AuthServiceClient
	serviceName string
	password    string
}

func NewAuthClient(cc *grpc.ClientConn, serviceName, password string) *AuthClient {
	service := pb.NewAuthServiceClient(cc)
	return &AuthClient{service, serviceName, password}
}

func (client *AuthClient) Login() (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &pb.LoginServiceRequest{
		ServiceName: client.serviceName,
		Password:    client.password,
	}

	res, err := client.service.Login(ctx, req)
	if err != nil {
		log.Println("Error logging in")
		return "", err
	}

	return res.GetAccessToken(), nil
}

func AuthMethods() map[string]bool {
	return map[string]bool{
		"/pb.AuthService/Login": true,
	}
}
