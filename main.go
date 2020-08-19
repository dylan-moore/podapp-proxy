package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var severCount = 0

const (
	SERVER1 = "https://getpod.app"
	PORT    = "1338"
)

func serveReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
	url, _ := url.Parse(target)

	proxy := httputil.NewSingleHostReverseProxy(url)

	proxy.ServeHTTP(res, req)
}

func logRequestPayload(proxyURL string) {
	log.Println("proxy_url: ", proxyURL)
}

func getProxyURL() string {
	var servers = []string{SERVER1}

	server := servers[severCount]
	severCount++

	if severCount >= len(servers) {
		severCount = 0
	}

	return server
}

func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {
	url := getProxyURL()

	log.Println("remoteAddr: ", req.RemoteAddr)

	logRequestPayload(url)

	serveReverseProxy(url, res, req)
}

func main() {
	log.Println("Began podapp-proxy")

	http.HandleFunc("*", handleRequestAndRedirect)

	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
