Introduction
==========
h2bench is a tool for testing how response size affects throughput when using HTTP/1.1 and/or HTTP/2. The tool is a minimalistic web server that only supports one request, `GET /random/{byteCount}` where `{byteCount}` is the number of bytes that will be returned from the web server. This can e.g. be used to measure the difference in throughput of 10000 requests that each returns 1k response versus 1000 requests each returning 100k response.

In order to test performance I recommend [h2load](https://nghttp2.org/documentation/h2load-howto.html).

Installation
============
You'll need [go][https://www.golang.org/] and [glide][https://github.com/Masterminds/glide] to build h2bench.
```
glide install
go build
```

Command line usage
==================
Usage of ./h2bench:
  -cert string
    	TLS certificate to use to secure the HTTP link.
  -key string
    	TLS private key to use to secure the HTTP link.
  -serverAddress string
    	The address to bind to (default ":8000")

Generating certificate
==================
You'll need a TLS certificate in order to use HTTP/2. To generate one, use:
```
openssl genrsa 2048 > server.key
chmod 400 server.key
openssl req -new -x509 -nodes -sha1 -days 3650 -key server.key > server.cert
```

Enter the address of your server when asked for FQDN.
