h1. To generate TLS certificate/key:

See https://gist.github.com/denji/12b3a568f092ab951456

openssl genrsa 2048 > server.key
chmod 400 server.key
openssl req -new -x509 -nodes -sha1 -days 3650 -key server.key > server.cert
