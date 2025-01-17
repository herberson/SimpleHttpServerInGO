package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello - now: %s\n", time.Now().String())
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func root(w http.ResponseWriter, req *http.Request) {
	headers(w, req)
	fmt.Fprintf(w, "\n\nhello\n\nnow: %s\n\n\n", time.Now().String())
}

func getLocalIPs() ([]net.IP, error) {
	var ips []net.IP
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	for _, addr := range addresses {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ips = append(ips, ipnet.IP)
			}
		}
	}
	return ips, nil
}

func main() {

	var port string

	if len(os.Args) > 1 {
		port = strings.TrimSpace(os.Args[1])
	} else {
		port = "8080"
	}

	ips, err := getLocalIPs()

	if err != nil {
		fmt.Println(err)
		ips = nil
	}

	fmt.Printf("\n=-=-=-= HTTP Server to routing test =-=-=-=\n\n")
	fmt.Printf("port: %s\n", port)
	fmt.Printf("   IP: %v\n\n", ips)

	fmt.Printf("Endpoints:\n")
	fmt.Printf("  /hello   - say's 'hello' and the current datetime of the server\n")
	fmt.Printf("  /headers - list request headers\n")
	fmt.Printf("  /        - same as '/hello'\n")
	fmt.Printf("\ncURL:\n")
	fmt.Printf("  curl -v http://%s:%s\n", ips[0], port)
	fmt.Printf("  curl -v http://%s:%s/hello\n", ips[0], port)
	fmt.Printf("  curl -v http://%s:%s/headers\n", ips[0], port)
	fmt.Printf("\n")

	port = ":" + port

	http.HandleFunc("/", root)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.ListenAndServe(port, nil)
}
