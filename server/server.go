package main

import (
	"fmt"
	"go-tls-demo-grpc/tlsservice"
)

func main() {
	fmt.Println("Hello World!")
	d := tlsservice.Data{
		D: "Some Message.",
	}
	fmt.Println(d)
}
