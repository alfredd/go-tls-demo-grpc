# USED FROM: https://dev.to/techschoolguru/how-to-secure-grpc-connection-with-ssl-tls-in-go-4ph

rm *.pem

# 1. Generate CA's private key and self-signed certificate
openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout ca-key.pem -out ca-cert.pem -subj "/C=CA/ST=Alberta/L=Edmonton/O=My CA/OU=Education/CN=*.alfredd.ca/emailAddress=ca@gmail.com"

echo "CA's self-signed certificate"
openssl x509 -in ca-cert.pem -noout -text

# 2. Generate web server's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout server-key.pem -out server-req.pem -subj "/C=CA/ST=Alberta/L=Edmonton/O=My Company/OU=Computer/CN=*.alfredd.ca/emailAddress=xyz@gmail.com"

# 3. Use CA's private key to sign web server's CSR and get back the signed certificate
openssl x509 -req -in server-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem -extfile server-ext.cnf
#Updated -extfile based on input from https://chowdera.com/2022/199/202207181303421208.html


echo "Server's signed certificate"
openssl x509 -in server-cert.pem -noout -text
