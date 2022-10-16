# Generate Protocol Buffer
`protoc --go_out=.  --go-grpc_out=.  proto/tlsservice.proto`

# Generate Certificates

```
$ cd cert
$ cert-gen.sh
```

# Prerequisites
1. Install Go gRPC compiler (https://grpc.io/docs/languages/go/quickstart/)

# Resources
1. https://dev.to/techschoolguru/how-to-secure-grpc-connection-with-ssl-tls-in-go-4ph
2. https://chowdera.com/2022/199/202207181303421208.html
