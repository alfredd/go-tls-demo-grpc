package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"go-tls-demo-grpc/tlsservice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"os"
	"time"
)

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	caFile := "./cert/ca-cert.pem"
	pemServerCA, err := os.ReadFile(caFile)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// Create the credentials and return it
	config := &tls.Config{
		RootCAs: certPool,
	}

	return credentials.NewTLS(config), nil
}
func main() {

	var opts []grpc.DialOption

	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}
	opts = append(opts, grpc.WithTransportCredentials(tlsCredentials))
	d := tlsservice.Data{D: "Client message"}
	conn, err := grpc.Dial("localhost:8081", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := tlsservice.NewTlsCommunicatorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := client.Message(ctx, &d)
	if err != nil {
		log.Fatalf("Error when sending message to server: %v", err)
	}
	fmt.Println(resp)
	time.Sleep(5 * time.Second)
}
