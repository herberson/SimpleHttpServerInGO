# Simple HTTP Server in GO

A very simple HTTP server in [GO](https://go.dev/) for network routing tests.

When I created this program, I couldn't use Python and if you arrived here because of the [code snippet](#python-simple-http-server), use it.

The program will only return a HTTP 200 response code with some output for "debug".

The intent is to help infrastructure teams on test the network routing configurations without the need of complex services or install Python on the target server.

## Using

This section will assume that you are using bash on a linux/unix distribution.

To start the server on 8080 port:

```sh
./SimpleHttpServerInGO_x86_lin
```

To choose a different port:

```sh
./SimpleHttpServerInGO_x86_lin 8088
```

The ouput on terminal for 8080 port is:

```
=-=-=-= HTTP Server to routing test =-=-=-=

port: 8080
   IP: [10.1.4.211 192.168.64.1 10.211.55.2 10.37.129.2]

Endpoints:
  /hello   - say's 'hello' and the current datetime of the server
  /headers - list request headers
  /        - same as '/hello'

cURL:
  curl -v http://10.1.4.211:8080
  curl -v http://10.1.4.211:8080/hello
  curl -v http://10.1.4.211:8080/headers

2025/01/17 18:14:38 URI      | Origin IP       | Remote Address  | HTTP Request Headers
```

When a request reach the program this generates a "log" as bellow:

```
2025/01/17 18:14:38 URI      | Origin IP       | Remote Address  | HTTP Request Headers
2025/01/17 18:14:43 /        | 127.0.0.1       | 127.0.0.1       | [User-Agent: curl/8.7.1] [Accept: */*]
2025/01/17 18:14:45 /hello   | 127.0.0.1       | 127.0.0.1       | [Accept: */*] [User-Agent: curl/8.7.1]
2025/01/17 18:14:47 /headers | 127.0.0.1       | 127.0.0.1       | [User-Agent: curl/8.7.1] [Accept: */*]
```

There you have:

* Requested URI
* Request origin IP, if present, show's ```X-Forwarded-For``` value.
* Request headers, delimited by brackets (```[]```).

The idea behind this log is an attempt to help on "debugging" (or "discover") the path witch the request had to follow to reach the target server.

## Build

This project is built with [Go](https://go.dev/) so to build your own binaries, you have to [install](https://go.dev/doc/install) and configure it as described on official documentation.

All code are on [main.go](main.go) file, so your modifications can be done on it.

To run your modifications on code do:

```sh
go run main.go
```

To build a executable run:

```sh
go build
```

With the current configuration the output file name is ```SimpleHttpServerInGO```, for windows it will have ```.exe``` extension.

The lack of comments on te code will be addressed on a later date. Contributions on this issue are welcomed.

## Python simple HTTP server

If you have Python available on your environment you can start a simple HTTP server with the command:

Python 3:
```sh
python3 -m http.server 8080
```

Python 2:
```sh
python -m SimpleHTTPServer 8080
```

If you arrived here because of the code snippet, use it.

## Conclusion

Hope this helps.