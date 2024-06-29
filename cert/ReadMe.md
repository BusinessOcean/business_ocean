The error message indicates that the TLS certificate verification failed because the certificate is not valid for the name localhost. This typically happens when the certificate does not include localhost in its Subject Alternative Names (SANs).

To resolve this issue, you need to generate a new TLS certificate that includes localhost in its SANs. Here are the steps to generate a new certificate with SANs using OpenSSL:


[req]
distinguished_name = req_distinguished_name
req_extensions = req_ext
x509_extensions = v3_req
prompt = no

[req_distinguished_name]
C = US
ST = California
L = San Francisco
O = MyCompany
OU = MyDivision
CN = localhost

[req_ext]
subjectAltName = @alt_names

[v3_req]
subjectAltName = @alt_names

[alt_names]
DNS.1 = localhost
IP.1 = 127.0.0.1


Create a configuration file (openssl.cnf) with the following content:
Generate a new private key and certificate using the configuration file:
Update your gRPC client code to use the new certificate:
This should resolve the certificate verification error by using a certificate that includes localhost in its SANs.


openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout server.key -out server.crt -config openssl.cnf