package main

import (
	"context"
	"fmt"
	"go-tls-demo-grpc/tlsservice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

type TlsServer struct {
	tlsservice.UnimplementedTlsCommunicatorServer
	clients []string
}

func (t *TlsServer) Message(ctx context.Context, data *tlsservice.Data) (*tlsservice.Data, error) {
	log.Println("Received message from client: ", data)
	return &tlsservice.Data{D: fmt.Sprintf("Response: %v", data.D)}, nil
}

func newTlsServer() *TlsServer {
	t := &TlsServer{}
	t.clients = make([]string, 1, 1)
	return t
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8081))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	certFile := "./cert/server-cert.pem"
	keyFile := "./cert/server-key.pem"

	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		log.Fatalf("Failed to generate credentials %v", err)
	}
	opts = []grpc.ServerOption{grpc.Creds(creds)}

	grpcServer := grpc.NewServer(opts...)
	tlsservice.RegisterTlsCommunicatorServer(grpcServer, newTlsServer())
	grpcServer.Serve(lis)
}
