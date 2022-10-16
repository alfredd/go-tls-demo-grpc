package main

import (
	"fmt"
	"go-tls-demo-grpc/tlsservice"
)

func main() {
	d := tlsservice.Data{D: "Client message"}
	fmt.Println(d)
}
