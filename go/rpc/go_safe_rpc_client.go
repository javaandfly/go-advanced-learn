package main

import (
	"fmt"
	"log"
	"net/rpc"
)

const HelloServiceName = "path/to/pkg.HelloService"

type HelloServiceClient struct {
	*rpc.Client
}

var _ HelloServiceInterface = (*HelloServiceClient)(nil)

type HelloServiceInterface interface {
	Hello(request string, reply *string) error
}

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

func main() {
	client, err := DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Hello("hello", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}

func (p *HelloServiceClient) Hello(request string, reply *string) error {
	err := p.Client.Call(HelloServiceName+".Hello", request, reply)
	return err
}
