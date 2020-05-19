package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"
)

var mapProxy map[string]*httputil.ReverseProxy

func initProxy() {
	// Initialize proxy with subdomains

	mapProxy = make(map[string]*httputil.ReverseProxy)
	mapPort := map[string]int{
		"www":   4000,
		"test":  4000,
		"map":   3000,
		"bbook": 4000}

	for subdomain, port := range mapPort {
		remote, _ := url.Parse("http://localhost:" + strconv.Itoa(port))
		mapProxy[subdomain] = httputil.NewSingleHostReverseProxy(remote)
	}
}

func subdomainProxy() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Ben", "Rad")

		// Check url contains subdomain
		subdomain := strings.Split(r.Host, ".")[0]

		mapProxy[subdomain].ServeHTTP(w, r)
	}
}

func main() {
	initProxy()
	
	log.Println("Reverse proxy start")

	http.HandleFunc("/", subdomainProxy())

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}
